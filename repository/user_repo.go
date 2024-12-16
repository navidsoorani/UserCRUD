package repository

import (
	"gorm.io/gorm"
	"simpleCrud/models"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uint) error
}

type GormUserRepository struct {
	DB *gorm.DB
}

// NewGormUserRepository creates a new instance of GormUserRepository
func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{DB: db}
}

// Create a new user
func (r *GormUserRepository) Create(user *models.User) (*models.User, error) {
	result := r.DB.Create(user)
	return user, result.Error
}

// Get all users
func (r *GormUserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	result := r.DB.Find(&users)
	return users, result.Error
}

// Get user by ID
func (r *GormUserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)
	return &user, result.Error
}

// Update a user
func (r *GormUserRepository) Update(user *models.User) (*models.User, error) {
	result := r.DB.Save(user)
	return user, result.Error
}

// Delete a user
func (r *GormUserRepository) Delete(id uint) error {
	result := r.DB.Delete(&models.User{}, id)
	return result.Error
}
