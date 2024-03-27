package repository

import (
	"BlogApp/domain"
	"context"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRatingRepository struct {
	database   *mongo.Database
	collection string
}

func NewBlogRatingRepository(db *mongo.Database, collection string) domain.BlogRatingRepository {
	return &BlogRatingRepository{
		database:   db,
		collection: collection,
	}
}

// DeleteRating implements domain.BlogRatingRepository.
func (b *BlogRatingRepository) DeleteRating(c context.Context, ratingID string) (*domain.BlogRating, error) {
	panic("unimplemented")
}

// GetRatingByBlogID implements domain.BlogRatingRepository.
func (b *BlogRatingRepository) GetRatingByBlogID(c context.Context, blogID string) (*[]*domain.BlogRating, error) {
	panic("unimplemented")
}

// GetRatingByUserID implements domain.BlogRatingRepository.
func (b *BlogRatingRepository) GetRatingByUserID(c context.Context, userID string) (*[]*domain.BlogRating, error) {
	panic("unimplemented")
}

// InsertRating implements domain.BlogRatingRepository.
func (b *BlogRatingRepository) InsertRating(c context.Context, rating *domain.BlogRating) (*domain.BlogRating, error) {
	panic("unimplemented")
}

// UpdateRating implements domain.BlogRatingRepository.
func (b *BlogRatingRepository) UpdateRating(c context.Context, rating *domain.BlogRating) (*domain.BlogRating, error) {
	panic("unimplemented")
}


