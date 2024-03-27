package repository

import (
	"BlogApp/domain"
	"context"
	"time"

	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	newFollow := &domain.Follow{
		FollowID:          primitive.NewObjectID().Hex(),
		FollowerID:        follow.FollowerID,
		FollowedID:        follow.FollowedID,
		CreatedAt:         time.Now(),

	  }
  
	  _, err := f.database.Collection(f.collection).InsertOne(c, newFollow)
	  if err != nil{ 
		return nil, err
	  }
	
	  return newFollow, nil
}

// Delete implements domain.FollowRepository.
func (f *FollowRepository) Delete(c context.Context, follow *domain.Follow) (*domain.Follow, error) {
	filter := bson.M{"_id": follow.FollowID}
	_, err := f.database.Collection(f.collection).DeleteOne(c, filter)
	if err != nil {
		return nil, err
	}
	return follow, nil

}

// GetByFollowedID implements domain.FollowRepository.
func (f *FollowRepository) GetByFollowedID(c context.Context, followedID string) (*[]*domain.Follow, error) {
	filter := bson.M{"followed_id": followedID}

	// Perform the find operation
	cursor, err := f.database.Collection(f.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	// Iterate through the cursor and decode each document into a User struct
	var follows []*domain.Follow
	for cursor.Next(c) {
		var follow domain.Follow
		if err := cursor.Decode(&follow); err != nil {
			cursor.Close(c)
			return nil, err
		}
		follows = append(follows, &follow)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &follows, nil
}

// GetByFollowerID implements domain.FollowRepository.
func (f *FollowRepository) GetByFollowerID(c context.Context, followerID string) (*[]*domain.Follow, error) {
	filter := bson.M{"follower_id": followerID}

	// Perform the find operation
	cursor, err := f.database.Collection(f.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	// Iterate through the cursor and decode each document into a User struct
	var follows []*domain.Follow
	for cursor.Next(c) {
		var follow domain.Follow
		if err := cursor.Decode(&follow); err != nil {
			cursor.Close(c)
			return nil, err
		}
		follows = append(follows, &follow)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &follows, nil
}


