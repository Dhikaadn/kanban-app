package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Where("user_id = ?", id).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	query := r.db.Create(category)
	if query.Error != nil {
		return 0, err
	}
	return category.ID, nil
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	for _, category := range categories {
		err := r.db.Create(&category).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	var user entity.Category
	query := r.db.First(&user, id)
	if query.Error != nil {
		return entity.Category{}, query.Error
	}
	return user, nil
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	query := r.db.Save(category)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	query := r.db.Delete(&entity.Category{}, id)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
