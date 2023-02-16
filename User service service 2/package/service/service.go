package service

import (
	"context"

	"github.com/Kin-dza-dzaa/testAssigment/internal/dto"
	"github.com/Kin-dza-dzaa/testAssigment/package/repository"
)

type Service interface {
	AddUser(context.Context, *dto.User) error
	GetUser(context.Context, string, *dto.UserDb) error
}

func NewService(repo repository.Repository) Service {
	return &ServiceUser{
		repo: repo,
	}
}
