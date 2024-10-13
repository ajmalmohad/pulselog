package repositories

import (
	"gorm.io/gorm"
)

type GenericRepository[T any] interface {
	Create(entity *T) (*T, error)
	Update(entity *T) (*T, error)
	Delete(entity *T) (*T, error)
	FindByID(id uint) (*T, error)
	FindAll() ([]*T, error)
}

type genericRepository[T any] struct {
	DB *gorm.DB
}

func NewGenericRepository[T any](db *gorm.DB) GenericRepository[T] {
	return &genericRepository[T]{DB: db}
}

func (r *genericRepository[T]) Create(entity *T) (*T, error) {
	result := r.DB.Create(entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

func (r *genericRepository[T]) Update(entity *T) (*T, error) {
	result := r.DB.Save(entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

func (r *genericRepository[T]) Delete(entity *T) (*T, error) {
	result := r.DB.Delete(entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

func (r *genericRepository[T]) FindByID(id uint) (*T, error) {
	var entity T
	result := r.DB.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

func (r *genericRepository[T]) FindAll() ([]*T, error) {
	var entities []*T
	result := r.DB.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}
