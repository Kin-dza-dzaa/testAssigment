package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Kin-dza-dza/testAssigment/internal/dto"
	"github.com/Kin-dza-dza/testAssigment/package/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Hanlers struct {
	Router  *mux.Router
	Service service.Service
}

func (h *Hanlers) CreateSaltHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		res := new(dto.SaltDto)
		res.Salt = h.Service.GenerateSalt()
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusOK)
	})
}

func (h *Hanlers) LoggingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logrus.WithFields(logrus.Fields{"url": r.RequestURI}).Info("serving request")
			next.ServeHTTP(w, r)
		})
	}
}

func NewHandlers(service service.Service) *Hanlers {
	handlers := new(Hanlers)
	handlers.Service = service
	handlers.Router = mux.NewRouter()
	handlers.Router.Handle("/generate-salt", handlers.CreateSaltHandler()).Methods("POST").Schemes("http")
	handlers.Router.Use(handlers.LoggingMiddleware())
	return handlers
}
