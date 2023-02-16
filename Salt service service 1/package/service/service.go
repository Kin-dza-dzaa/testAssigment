package service

type Service interface {
	GenerateSalt() string
}

func NewService() Service {
	return &SaltService{}
}
