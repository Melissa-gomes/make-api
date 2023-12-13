package services

import (
	"project-api/repositories"
)

type Service struct {
	Repo repositories.Repository
}

func NewService(repo repositories.Repository) Service {
	return Service{Repo: repo}
}

func (s Service) ServiceGet() string {
	return "oi"
}
