package repository

import "bootcamp_auth_microservice/infras"

type Repository interface {
	UserRepository
}

type RepositoryImpl struct {
	DB *infras.Conn
}

func ProvideRepo(db *infras.Conn) *RepositoryImpl {
	return &RepositoryImpl{
		DB: db,
	}
}
