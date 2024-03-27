package repository

import (
	"BlogApp/domain"
	"context"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepository struct {
	database   *mongo.Database
	collection string
}

func NewCommentRepository(db *mongo.Database, collection string) domain.CommentRepository {
	return &CommentRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.CommentRepository.
func (*CommentRepository) Create(c context.Context, comment *domain.Comment) (*domain.Comment, error) {
	
}

// Delete implements domain.CommentRepository.
func (*CommentRepository) Delete(c context.Context, commentID string) (*domain.Comment, error) {
	panic("unimplemented")
}

// GetByBlogID implements domain.CommentRepository.
func (*CommentRepository) GetByBlogID(c context.Context, blogID string) (*[]*domain.Comment, error) {
	panic("unimplemented")
}

// Update implements domain.CommentRepository.
func (*CommentRepository) Update(c context.Context, comment *domain.Comment) (*domain.Comment, error) {
	panic("unimplemented")
}


