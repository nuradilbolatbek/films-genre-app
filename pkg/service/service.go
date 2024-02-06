package service

import (
	"FilmsProject"
	"FilmsProject/pkg/repository"
)

type Authorization interface {
	CreateUser(user FilmsProject.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(AccessToken string) (int, error)
}

type GenreList interface {
	Create(userId int, list FilmsProject.Genrelist) (int, error)
	GetAll(userId int) ([]FilmsProject.Genrelist, error)
	GetById(userId int, listId int) (FilmsProject.Genrelist, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input FilmsProject.UpdateList) error
}

type GenreFilms interface {
	Create(userId, listId int, item FilmsProject.GenreFilms) (int, error)
	GetAll(userId, listId int) ([]FilmsProject.GenreFilms, error)
	GetById(userId int, itemId int) (FilmsProject.GenreFilms, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input FilmsProject.UpdateItem) error
}
type Service struct {
	Authorization
	GenreFilms
	GenreList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		GenreList:     NewGenreListService(repos.GenreList),
		GenreFilms:    NewGenreFilmsService(repos.GenreFilms, repos.GenreList),
	}
}
