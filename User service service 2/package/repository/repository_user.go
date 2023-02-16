package repository

import (
	"context"

	"github.com/Kin-dza-dzaa/testAssigment/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryUser struct {
	collection *mongo.Collection
}

func (r *RepositoryUser) IfUserExists(ctx context.Context, email string) bool {
	return r.collection.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Err() != mongo.ErrNoDocuments
}

func (r *RepositoryUser) AddUser(ctx context.Context, user *dto.UserDb) error {
	if _, err := r.collection.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}

func (r *RepositoryUser) GetUser(ctx context.Context, email string, dbUser *dto.UserDb) error {
	res := r.collection.FindOne(ctx, bson.D{{Key: "email", Value: email}})
	if err := res.Err(); err != nil {
		return err
	}
	return res.Decode(dbUser)
}
