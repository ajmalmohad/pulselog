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
	db *gorm.DB
}

func NewGenericRepository[T any](db *gorm.DB) GenericRepository[T] {
	return &genericRepository[T]{db: db}
}

func (r *genericRepository[T]) Create(entity *T) (*T, error) {
	result := r.db.Create(entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

func (r *genericRepository[T]) Update(entity *T) (*T, error) {
	result := r.db.Save(entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

func (r *genericRepository[T]) Delete(entity *T) (*T, error) {
	result := r.db.Delete(entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

func (r *genericRepository[T]) FindByID(id uint) (*T, error) {
	var entity T
	result := r.db.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

func (r *genericRepository[T]) FindAll() ([]*T, error) {
	var entities []*T
	result := r.db.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}
