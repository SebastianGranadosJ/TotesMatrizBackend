package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type PermissionService struct {
	Repo *repositories.PermissionRepository
}

func NewPermissionService(repo *repositories.PermissionRepository) *PermissionService {
	return &PermissionService{Repo: repo}
}

func (s *PermissionService) GetPermissionByID(id uint) (*models.Permission, error) {
	return s.Repo.GetPermissionByID(id)
}

func (s *PermissionService) GetAllPermissions() ([]models.Permission, error) {
	return s.Repo.GetAllPermissions()
}
