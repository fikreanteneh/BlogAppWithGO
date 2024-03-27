package repository

import (
	"BlogApp/domain"
	"context"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository struct {
	database   *mongo.Database
	collection string
}

func NewBlogRepository(db *mongo.Database, collection string) domain.BlogRepository {
	return &BlogRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.BlogRepository.
func (b *BlogRepository) Create(c context.Context, blog *domain.Blog) (*domain.Blog, error) {
	panic("unimplemented")
}

// Delete implements domain.BlogRepository.
func (b *BlogRepository) Delete(c context.Context, blogID string) (*domain.Blog, error) {
	panic("unimplemented")
}

// GetAll implements domain.BlogRepository.
func (b *BlogRepository) GetAll(c context.Context, param string) (*[]*domain.Blog, error) {
	panic("unimplemented")
}

// GetByID implements domain.BlogRepository.
func (b *BlogRepository) GetByID(c context.Context, blogID string) (*domain.Blog, error) {
	panic("unimplemented")
}

// GetByUserId implements domain.BlogRepository.
func (b *BlogRepository) GetByUserId(c context.Context, userID string) (*[]*domain.Blog, error) {
	panic("unimplemented")
}

// Update implements domain.BlogRepository.
func (b *BlogRepository) Update(c context.Context, blog *domain.Blog) (*domain.Blog, error) {
	panic("unimplemented")
}


