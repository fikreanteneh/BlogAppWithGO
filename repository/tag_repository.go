package repository

import (
	"BlogApp/domain"
	"context"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagRepository struct {
	database   *mongo.Database
	collection string
}

func NewTagRepository(db *mongo.Database, collection string) domain.TagRepository {
	return &TagRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.TagRepository.
func (t *TagRepository) Create(c context.Context, tag *domain.Tag) (*domain.Tag, error) {
	panic("unimplemented")
}

// Delete implements domain.TagRepository.
func (t *TagRepository) Delete(c context.Context, tagID string) (*domain.Tag, error) {
	panic("unimplemented")
}

// GetAll implements domain.TagRepository.
func (t *TagRepository) GetAll(c context.Context, param string) (*[]*domain.Tag, error) {
	panic("unimplemented")
}

// GetByID implements domain.TagRepository.
func (t *TagRepository) GetByID(c context.Context, tagID string) (*domain.Tag, error) {
	panic("unimplemented")
}

// GetByName implements domain.TagRepository.
func (t *TagRepository) GetByName(c context.Context, name string) (*domain.Tag, error) {
	panic("unimplemented")
}

// Update implements domain.TagRepository.
func (t *TagRepository) Update(c context.Context, tag *domain.Tag) (*domain.Tag, error) {
	panic("unimplemented")
}


