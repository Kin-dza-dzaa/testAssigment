package hanlders

import (
	"encoding/json"
	"net/http"

	"github.com/Kin-dza-dzaa/testAssigment/internal/dto"
	customlog "github.com/Kin-dza-dzaa/testAssigment/internal/structure_logger"
	"github.com/Kin-dza-dzaa/testAssigment/package/service"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	service service.Service
	Router  chi.Router
}

func (h *Handlers) GetUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		user := new(dto.UserDb)
		if err := h.service.GetUser(r.Context(), chi.URLParam(r, "email"), user); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusOK)
	})
}

func (h *Handlers) CreateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		user := new(dto.User)
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			json.NewEncoder(w).Encode(map[string]string{"result": "error", "expected": "email, password"})
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.service.AddUser(r.Context(), user); err != nil {
			logrus.Error(err)
			if err == service.ErrInvalidEmail {
				json.NewEncoder(w).Encode(map[string]string{"result": "error", "message": "invalid email"})
			}
			if err == service.ErrUserAlreadyExists {
				json.NewEncoder(w).Encode(map[string]string{"result": "error", "message": "user already exists"})
			}
			if err == service.ErrSaltServiceIsDown {
				json.NewEncoder(w).Encode(map[string]string{"result": "error", "message": "salt server is down"})
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func NewHandlers(serv service.Service) *Handlers {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}
	handlers := new(Handlers)
	handlers.service = serv
	handlers.Router = chi.NewRouter()
	handlers.Router.Use(customlog.NewStructuredLogger(logger))
	handlers.Router.Post("/create-user", handlers.CreateUser().ServeHTTP)
	handlers.Router.Get("/get-user/{email}", handlers.GetUser().ServeHTTP)
	return handlers
}
