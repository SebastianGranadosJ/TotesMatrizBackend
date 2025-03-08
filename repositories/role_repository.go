package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (r *RoleRepository) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	err := r.DB.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepository) GetRoleByID(id uint) (*models.Role, error) {
	var role models.Role
	err := r.DB.Preload("Permissions").First(&role, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) GetRolePermissions(roleID uint) ([]uint, error) {
	var permissionIDs []uint
	err := r.DB.Table("role_permission").Select("permission_id").
		Where("role_id = ?", roleID).
		Pluck("permission_id", &permissionIDs).Error
	if err != nil {
		return nil, err
	}
	return permissionIDs, nil
}

func (r *RoleRepository) GetAllPermissionsOfRole(roleID uint) ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.DB.Joins("JOIN role_permission rp ON permissions.id = rp.permission_id").
		Where("rp.role_id = ?", roleID).
		Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *RoleRepository) ExistRole(roleID uint) (bool, error) {
	var count int64
	err := r.DB.Table("roles").Where("id = ?", roleID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
