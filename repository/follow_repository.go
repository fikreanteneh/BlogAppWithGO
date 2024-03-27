package repository

import (
	"BlogApp/domain"
	"context"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FollowRepository struct {
	database   *mongo.Database
	collection string
}

func NewFollowRepository(db *mongo.Database, collection string) domain.FollowRepository {
	return &FollowRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.FollowRepository.
func (f *FollowRepository) Create(c context.Context, follow *domain.Follow) (*domain.Follow, error) {
	panic("unimplemented")
}

// Delete implements domain.FollowRepository.
func (f *FollowRepository) Delete(c context.Context, follow *domain.Follow) (*domain.Follow, error) {
	panic("unimplemented")
}

// GetByFollowedID implements domain.FollowRepository.
func (f *FollowRepository) GetByFollowedID(c context.Context, followedID string) (*[]*domain.Follow, error) {
	panic("unimplemented")
}

// GetByFollowerID implements domain.FollowRepository.
func (f *FollowRepository) GetByFollowerID(c context.Context, followerID string) (*[]*domain.Follow, error) {
	panic("unimplemented")
}


