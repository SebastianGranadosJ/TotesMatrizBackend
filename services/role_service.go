package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type RoleService struct {
	Repo *repositories.RoleRepository
}

func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{Repo: repo}
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	return s.Repo.GetAllRoles()
}

func (s *RoleService) GetRoleByID(id uint) (*models.Role, error) {
	return s.Repo.GetRoleByID(id)

}

func (s *RoleService) GetRolePermissions(roleID uint) ([]uint, error) {
	return s.Repo.GetRolePermissions(roleID)
}

func (s *RoleService) GetAllPermissionsOfRole(roleID uint) ([]models.Permission, error) {
	return s.Repo.GetAllPermissionsOfRole(roleID)
}

func (s *RoleService) ExistRole(roleID uint) (bool, error) {
	return s.Repo.ExistRole(roleID)
}
