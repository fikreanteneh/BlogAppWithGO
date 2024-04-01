package repository

import (
	"BlogApp/domain"
	"context"

	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShareRepository struct {
	database   *mongo.Database
	collection string
}

// DeleteShareByBlogID implements domain.ShareRepository.
func (s *ShareRepository) DeleteShareByBlogID(c context.Context, blogID string) error {
	filter := bson.M{"blog_id": blogID}

	_, err := s.database.Collection(s.collection).DeleteMany(c, filter)
	if err != nil {
		return err
	}
	return nil
}

func NewShareRepository(db *mongo.Database, collection string) domain.ShareRepository {
	return &ShareRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.ShareRepository.
func (s *ShareRepository) Create(c context.Context, share *domain.Share) (*domain.Share, error) {
	share.ShareID = primitive.NewObjectID().Hex()

	_, err := s.database.Collection(s.collection).InsertOne(c, *share)
	if err != nil {
		return nil, err
	}

	return share, nil
}

// Delete implements domain.ShareRepository.
func (s *ShareRepository) Delete(c context.Context, shareID string) (*domain.Share, error) {
	filter := bson.M{"_id": shareID}

	var share domain.Share
	err := s.database.Collection(s.collection).FindOneAndDelete(c, filter).Decode(&share)
	if err != nil {
		return nil, err
	}

	return &share, nil

}

// GetByBlogID implements domain.ShareRepository.
func (s *ShareRepository) GetByBlogID(c context.Context, blogID string) (*[]*domain.Share, error) {
	filter := bson.M{"blog_id": blogID}

	cursor, err := s.database.Collection(s.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	// Iterate through the cursor and decode each document into a Like struct
	var shares []*domain.Share
	for cursor.Next(c) {
		var share domain.Share
		if err := cursor.Decode(&share); err != nil {
			cursor.Close(c)
			return nil, err
		}
		shares = append(shares, &share)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &shares, nil
}

// GetByUserID implements domain.ShareRepository.
func (s *ShareRepository) GetByUserID(c context.Context, userID string) (*[]*domain.Share, error) {
	filter := bson.M{"user_id": userID}

	cursor, err := s.database.Collection(s.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	// Iterate through the cursor and decode each document into a Like struct
	var shares []*domain.Share
	for cursor.Next(c) {
		var share domain.Share
		if err := cursor.Decode(&share); err != nil {
			cursor.Close(c)
			return nil, err
		}
		shares = append(shares, &share)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &shares, nil
}
