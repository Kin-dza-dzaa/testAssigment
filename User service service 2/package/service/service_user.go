package service

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/Kin-dza-dzaa/testAssigment/internal/dto"
	"github.com/Kin-dza-dzaa/testAssigment/package/repository"
	"github.com/sirupsen/logrus"
)

var (
	ErrInvalidEmail      = errors.New("invalid email")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrSaltServiceIsDown = errors.New("salt server is down")
)

type ServiceUser struct {
	repo repository.Repository
}

func (s *ServiceUser) GetUser(ctx context.Context, email string, user *dto.UserDb) error {
	return s.repo.GetUser(ctx, email, user)
}

func (s *ServiceUser) AddUser(ctx context.Context, user *dto.User) error {
	if !s.validateEmailUsingRegex(user.Email) {
		return ErrInvalidEmail
	}
	if s.repo.IfUserExists(ctx, user.Email) {
		return ErrUserAlreadyExists
	}
	salt, err := s.getSaltFromService()
	if err != nil {
		return err
	}

	hashedPassword := s.hashPassword(user.Password, salt)
	userDb := new(dto.UserDb)
	userDb.Email = user.Email
	userDb.Password = hashedPassword
	userDb.Salt = salt

	return s.repo.AddUser(ctx, userDb)
}

func (s *ServiceUser) validateEmailUsingRegex(email string) bool {
	defer func() {
		if err := recover(); err != nil {
			logrus.WithFields(logrus.Fields{
				"event": "validateEmailUsingRegex",
			}).Warnf("recovered with err: %v", err)
		}
	}()
	expr := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return expr.MatchString(email)
}

func (s *ServiceUser) getSaltFromService() (string, error) {
	res, err := http.Post(dto.Cfg.SaltServiceAddress, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	salt := new(dto.Salt)
	if err := json.NewDecoder(res.Body).Decode(salt); err != nil {
		return "", err
	}

	return salt.Salt, nil
}

func (s *ServiceUser) hashPassword(password string, salt string) string {
	h := md5.New()
	io.WriteString(h, password)
	io.WriteString(h, salt)
	return fmt.Sprintf("%x", h.Sum(nil))
}
