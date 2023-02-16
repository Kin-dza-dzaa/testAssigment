package main

import (
	"net/http"
	"time"

	"github.com/Kin-dza-dza/testAssigment/internal/dto"
	"github.com/Kin-dza-dza/testAssigment/package/handlers"
	"github.com/Kin-dza-dza/testAssigment/package/service"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	if err := dto.ReadConfig(); err != nil {
		logrus.WithFields(logrus.Fields{
			"event": "init logger",
		}).Fatal(err)
	}
}

func main() {
	handlers := handlers.NewHandlers(service.NewService())
	srv := &http.Server{
		Addr:         dto.Cfg.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      handlers.Router,
	}
	go func() {
		logrus.WithFields(logrus.Fields{
			"event": "start server",
		}).Infof("Salt server is ready to accept calls` at %v", dto.Cfg.Address)
		if err := srv.ListenAndServe(); err != nil {
			logrus.WithFields(logrus.Fields{
				"event": "start server",
			}).Fatal(err)
		}
	}()
	select {}
}
