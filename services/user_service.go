package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) SearchUsersByID(query string) ([]models.User, error) {
	return s.Repo.SearchUsersByID(query)
}

func (s *UserService) SearchUsersByEmail(query string) ([]models.User, error) {
	return s.Repo.SearchUsersByEmail(query)
}

func (s *UserService) UpdateUserState(id string, state int) (*models.User, error) {
	return s.Repo.UpdateUserState(id, state)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.Repo.UpdateUser(user)
}
