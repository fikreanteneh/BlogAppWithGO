package repository

import (
	"BlogApp/domain"
	"context"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository struct {
	database   *mongo.Database
	collection string
}

func NewLikeRepository(db *mongo.Database, collection string) domain.LikeRepository {
	return &LikeRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.LikeRepository.
func (l *LikeRepository) Create(c context.Context, like *domain.Like) (*domain.Like, error) {
	panic("unimplemented")
}

// Delete implements domain.LikeRepository.
func (l *LikeRepository) Delete(c context.Context, likeID string) (*domain.Like, error) {
	panic("unimplemented")
}

// GetByBlogID implements domain.LikeRepository.
func (l *LikeRepository) GetByBlogID(c context.Context, blogID string) (*[]*domain.Like, error) {
	panic("unimplemented")
}

// GetByUserID implements domain.LikeRepository.
func (l *LikeRepository) GetByUserID(c context.Context, userID string) (*[]*domain.Like, error) {
	panic("unimplemented")
}


