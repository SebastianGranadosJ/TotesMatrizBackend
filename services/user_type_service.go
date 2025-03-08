package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type UserTypeService struct {
	Repo *repositories.UserTypeRepository
}

func NewUserTypeService(repo *repositories.UserTypeRepository) *UserTypeService {
	return &UserTypeService{Repo: repo}
}

func (s *UserTypeService) ObtainAllUserTypes() ([]models.UserType, error) {
	return s.Repo.ObtainAllUserTypes()
}

func (s *UserTypeService) ObtainUserTypeByID(id uint) (*models.UserType, error) {
	return s.Repo.ObtainUserTypeByID(id)
}

func (s *UserTypeService) GetRolesForUserType(userTypeID uint) ([]uint, error) {
	return s.Repo.GetRolesForUserType(userTypeID)
}

func (s *UserTypeService) Exists(userTypeID uint) (bool, error) {
	return s.Repo.Exists(userTypeID)
}
