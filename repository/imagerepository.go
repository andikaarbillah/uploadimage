package repository

import (
	"images/model"

	"gorm.io/gorm"
)

type ImageRepository interface {
	Create(image model.Images) (*model.Images, error)
	Delete(imageID string) error
	FindByID(imageID string) (*model.Images, error)
}

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) ImageRepository {
	return &imageRepository{
		db: db,
	}
}

func (ir *imageRepository) Create(image model.Images) (*model.Images, error) {
	result := ir.db.Create(&image)
	if result.Error != nil {
		return nil, result.Error
	}

	return &image, nil
}

func (ir *imageRepository) Delete(imageID string) error {
	result := ir.db.Delete(&model.Images{}, imageID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ir *imageRepository) FindByID(imageID string) (*model.Images, error) {
	var image model.Images
	result := ir.db.First(&image, "id = ?", imageID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &image, nil
}
