package service

import (
	"FilmsProject"
	"FilmsProject/pkg/repository"
)

type GenrelistService struct {
	repo repository.GenreList
}

func NewGenreListService(repo repository.GenreList) *GenrelistService {
	return &GenrelistService{repo: repo}
}

func (s *GenrelistService) Create(userId int, list FilmsProject.Genrelist) (int, error) {
	return s.repo.Create(userId, list)
}
func (s *GenrelistService) GetAll(userId int) ([]FilmsProject.Genrelist, error) {
	return s.repo.GetAll(userId)
}
func (s *GenrelistService) GetById(userId int, listId int) (FilmsProject.Genrelist, error) {
	return s.repo.GetById(userId, listId)
}

func (s *GenrelistService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *GenrelistService) Update(userId, listId int, input FilmsProject.UpdateList) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
