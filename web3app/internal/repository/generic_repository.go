package repository

import (
	"errors"

	"gorm.io/gorm"
)

type Repository[T any] interface {
	Insert(entity T) (T, error)
	SelectAll() ([]T, error)
	SelectById(id string) (T, error)
	Update(entity T) (T, error)
	UpdateById(id string, updates map[string]interface{}) (T, error)
	SelectByField(field string, value interface{}) (T, error)
	SelectByRequestId(id string) (T, error)
	SelectByUID(uid string) (T, error)
	ShowByField(field string, value interface{}) ([]T, error)

	UpdateByRequestId(id string, updates map[string]interface{}) (T, error)
	UpdateByUID(id string, updates map[string]interface{}) (T, error)
}

// repository struct implements the Repository interface.
type repository[T any] struct {
	db *gorm.DB
}

// NewRepository creates a new repository.
func NewRepository[T any](db *gorm.DB) Repository[T] {
	return &repository[T]{db: db}
}

// Insert inserts a new entity into the database.
func (r *repository[T]) Insert(entity T) (T, error) {
	err := r.db.Create(&entity).Error
	return entity, err
}

// SelectAll retrieves all entities from the database.
func (r *repository[T]) SelectAll() ([]T, error) {
	var entities []T
	err := r.db.Order("created_at asc").Find(&entities).Error
	return entities, err
}

// SelectById retrieves an entity by its ID from the database.
func (r *repository[T]) SelectById(id string) (T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	return entity, err
}

func (r *repository[T]) SelectByRequestId(id string) (T, error) {
	var entity T
	err := r.db.First(&entity, "request_id = ?", id).Error
	return entity, err
}

// SelectById retrieves an entity by its UID from the database.
func (r *repository[T]) SelectByUID(uid string) (T, error) {
	var entity T
	err := r.db.Where("uid = ?", uid).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity, nil // or handle not found case as needed
	}
	return entity, err
}

// SelectByName retrieves an entity by a specific field and value from the database.
func (r *repository[T]) SelectByField(field string, value interface{}) (T, error) {
	var entity T
	err := r.db.Where(field+" = ?", value).First(&entity).Error
	return entity, err
}

// // SelectByName retrieves an entity by a specific field and value from the database.
func (r *repository[T]) ShowByField(field string, value interface{}) ([]T, error) {
	var entity []T
	err := r.db.Where(field+" = ?", value).Find(&entity).Error
	return entity, err
}

// Update updates an existing entity in the database.
func (r *repository[T]) Update(entity T) (T, error) {
	err := r.db.Save(&entity).Error
	return entity, err
}

// Update updates the specified fields of an existing entity in the database.
func (r *repository[T]) UpdateById(id string, updates map[string]interface{}) (T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	if err != nil {
		return entity, err
	}

	err = r.db.Model(&entity).Updates(updates).Error
	return entity, err
}

func (r *repository[T]) UpdateByRequestId(id string, updates map[string]interface{}) (T, error) {
	var entity T
	err := r.db.First(&entity, "request_id = ?", id).Error
	if err != nil {
		return entity, err
	}

	err = r.db.Model(&entity).Updates(updates).Error
	return entity, err
}

func (r *repository[T]) UpdateByUID(id string, updates map[string]interface{}) (T, error) {
	var entity T
	err := r.db.First(&entity, "uid = ?", id).Error
	if err != nil {
		return entity, err
	}

	err = r.db.Model(&entity).Updates(updates).Error
	return entity, err
}
