package main

import (
	"context"
	"net/http"
	"time"

	"github.com/Kin-dza-dzaa/testAssigment/internal/dto"
	hanlders "github.com/Kin-dza-dzaa/testAssigment/package/handlers"
	"github.com/Kin-dza-dzaa/testAssigment/package/repository"
	"github.com/Kin-dza-dzaa/testAssigment/package/service"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dto.Cfg.DbUrl))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event": "mongodb connect",
		}).Fatal(err)
	}

	repo := repository.NewRepository(client.Database(dto.Cfg.DbName).Collection(dto.Cfg.CollectionName))
	serv := service.NewService(repo)
	hand := hanlders.NewHandlers(serv)
	srv := &http.Server{
		Addr:         dto.Cfg.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      hand.Router,
	}
	go func() {
		logrus.Infof("Starting salt server at %v", dto.Cfg.Address)
		if err := srv.ListenAndServe(); err != nil {
			logrus.WithFields(logrus.Fields{
				"event": "start server",
			}).Fatal(err)
		}
	}()
	select {}
}
