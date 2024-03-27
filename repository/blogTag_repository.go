package repository

import (
	"BlogApp/domain"
	"context"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogTagRepository struct {
	database   *mongo.Database
	collection string
}

func NewBlogTagRepository(db *mongo.Database, collection string) domain.BlogTagRepository {
	return &BlogTagRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.BlogTagRepository.
func (b *BlogTagRepository) Create(c context.Context, blogTag *domain.BlogTag) (*domain.BlogTag, error) {
	panic("unimplemented")
}

// Delete implements domain.BlogTagRepository.
func (b *BlogTagRepository) Delete(c context.Context, blogTagID string) (*domain.BlogTag, error) {
	panic("unimplemented")
}

// GetByBlogID implements domain.BlogTagRepository.
func (b *BlogTagRepository) GetByBlogID(c context.Context, blogID string) (*[]*domain.BlogTag, error) {
	panic("unimplemented")
}

// GetByTagID implements domain.BlogTagRepository.
func (b *BlogTagRepository) GetByTagID(c context.Context, tagID string) (*[]*domain.BlogTag, error) {
	panic("unimplemented")
}


