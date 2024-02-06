package service

import (
	"FilmsProject"
	"FilmsProject/pkg/repository"
)

type GenreFilmsService struct {
	repo     repository.GenreFilms
	listRepo repository.GenreList
}

func NewGenreFilmsService(repo repository.GenreFilms, listRepo repository.GenreList) *GenreFilmsService {
	return &GenreFilmsService{repo: repo, listRepo: listRepo}
}

func (s *GenreFilmsService) Create(userId, listId int, item FilmsProject.GenreFilms) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, nil
	}
	return s.repo.Create(listId, item)
}

func (s *GenreFilmsService) GetAll(userId, listId int) ([]FilmsProject.GenreFilms, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *GenreFilmsService) GetById(userId, itemId int) (FilmsProject.GenreFilms, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *GenreFilmsService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *GenreFilmsService) Update(userId, itemId int, input FilmsProject.UpdateItem) error {
	return s.repo.Update(userId, itemId, input)
}
