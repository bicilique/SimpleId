package repository

import (
	"SimpleId/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	SelectAll() ([]entity.User, error)
	Insert(entity entity.User) (entity.User, error)
	SelectById(id string) (entity.User, error)
	SelectByName(field string, value interface{}) (entity.User, error)
	Update(entity entity.User) (entity.User, error)
	UpdateById(id string, updates map[string]interface{}) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewRepository creates a new repository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Insert inserts a new entity into the database.
func (r *userRepository) Insert(entity entity.User) (entity.User, error) {
	err := r.db.Create(&entity).Error
	return entity, err
}

// SelectAll retrieves all entities from the database.
func (r *userRepository) SelectAll() ([]entity.User, error) {
	var entities []entity.User
	err := r.db.Find(&entities).Error
	return entities, err
}

// SelectById retrieves an entity by its ID from the database.
func (r *userRepository) SelectById(id string) (entity.User, error) {
	var entity entity.User
	err := r.db.First(&entity, id).Error
	return entity, err
}

// SelectByName retrieves an entity by a specific field and value from the database.
func (r *userRepository) SelectByName(field string, value interface{}) (entity.User, error) {
	var entity entity.User
	err := r.db.Where(field+" = ?", value).First(&entity).Error
	return entity, err
}

// Update updates an existing entity in the database.
func (r *userRepository) Update(entity entity.User) (entity.User, error) {
	err := r.db.Save(&entity).Error
	return entity, err
}

// Update updates the specified fields of an existing entity in the database.
func (r *userRepository) UpdateById(id string, updates map[string]interface{}) (entity.User, error) {
	var entity entity.User
	err := r.db.First(&entity, id).Error
	if err != nil {
		return entity, err
	}

	err = r.db.Model(&entity).Updates(updates).Error
	return entity, err
}
