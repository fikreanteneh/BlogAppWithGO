package repository

import (
	"BlogApp/domain"
	"context"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShareRepository struct {
	database   *mongo.Database
	collection string
}

func NewShareRepository(db *mongo.Database, collection string) domain.ShareRepository {
	return &ShareRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.ShareRepository.
func (s *ShareRepository) Create(c context.Context, share *domain.Share) (*domain.Share, error) {
	panic("unimplemented")
}

// Delete implements domain.ShareRepository.
func (s *ShareRepository) Delete(c context.Context, shareID string) (*domain.Share, error) {
	panic("unimplemented")
}

// GetByBlogID implements domain.ShareRepository.
func (s *ShareRepository) GetByBlogID(c context.Context, blogID string) (*[]*domain.Share, error) {
	panic("unimplemented")
}

// GetByUserID implements domain.ShareRepository.
func (s *ShareRepository) GetByUserID(c context.Context, userID string) (*[]*domain.Share, error) {
	panic("unimplemented")
}


