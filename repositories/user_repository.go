package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.DB.Preload("UserStateType").Preload("UserType").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Preload("UserStateType").Preload("UserType").First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Preload("UserStateType").Preload("UserType").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) SearchUsersByID(query string) ([]models.User, error) {
	var users []models.User
	err := r.DB.Preload("UserStateType").Preload("UserType").
		Where("CAST(id AS TEXT)  LIKE ?", query+"%").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) SearchUsersByEmail(query string) ([]models.User, error) {
	var users []models.User
	err := r.DB.Preload("UserStateType").Preload("UserType").Where("LOWER(email) LIKE LOWER(?)", query+"%").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) UpdateUserState(id string, state int) (*models.User, error) {
	var user models.User
	if err := r.DB.Preload("UserStateType").Preload("UserType").First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	user.UserStateType.ID = state

	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	var existingUser models.User
	if err := r.DB.Preload("UserStateType").Preload("UserType").First(&existingUser, "id = ?", user.ID).Error; err != nil {
		return err
	}
	// Realizar la actualización
	if err := r.DB.Model(&existingUser).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	// Intentar crear el usuario en la base de datos
	if err := r.DB.Preload("UserStateType").Preload("UserType").Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
