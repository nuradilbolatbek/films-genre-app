package repository

import (
	"FilmsProject"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user FilmsProject.User) (int, error)
	GetUser(username, password string) (FilmsProject.User, error)
}

type GenreList interface {
	Create(userId int, list FilmsProject.Genrelist) (int, error)
	GetAll(userId int) ([]FilmsProject.Genrelist, error)
	GetById(userId int, listId int) (FilmsProject.Genrelist, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input FilmsProject.UpdateList) error
}

type GenreFilms interface {
	Create(listId int, item FilmsProject.GenreFilms) (int, error)
	GetAll(userId, listId int) ([]FilmsProject.GenreFilms, error)
	GetById(userId int, itemId int) (FilmsProject.GenreFilms, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input FilmsProject.UpdateItem) error
}
type Repository struct {
	Authorization
	GenreFilms
	GenreList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		GenreList:     NewGenreListPostgres(db),
		GenreFilms:    NewGenreFilmsPostgres(db),
	}

}
