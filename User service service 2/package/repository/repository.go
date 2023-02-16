package repository

import (
	"context"

	"github.com/Kin-dza-dzaa/testAssigment/internal/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	AddUser(context.Context, *dto.UserDb) error
	IfUserExists(context.Context, string) bool
	GetUser(context.Context, string, *dto.UserDb) error
}

func NewRepository(collection *mongo.Collection) Repository {
	return &RepositoryUser{
		collection: collection,
	}
}
