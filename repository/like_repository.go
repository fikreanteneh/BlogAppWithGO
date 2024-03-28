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
	like.LikeID = primitive.NewObjectID().Hex()
	  _, err := l.database.Collection(l.collection).InsertOne(c, *like)
	  if err != nil{ 
		return nil, err
	  }
	
	  return like, nil
	
}

// Delete implements domain.LikeRepository.
func (l *LikeRepository) Delete(c context.Context, likeID string) (*domain.Like, error) {
	filter := bson.M{"_id": likeID}

	var like domain.Like
	err := l.database.Collection(l.collection).FindOneAndDelete(c, filter).Decode(&like)
	if err != nil {
		return nil, err
	}

	return &like, nil

}

// GetByBlogID implements domain.LikeRepository.
func (l *LikeRepository) GetByBlogID(c context.Context, blogID string) (*[]*domain.Like, error) {
	filter := bson.M{"blog_id": blogID}

	cursor, err := l.database.Collection(l.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	// Iterate through the cursor and decode each document into a Like struct
	var likes []*domain.Like
	for cursor.Next(c) {
		var like domain.Like
		if err := cursor.Decode(&like); err != nil {
			cursor.Close(c)
			return nil, err
		}
		likes = append(likes, &like)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &likes, nil
}

// GetByUserID implements domain.LikeRepository.
func (l *LikeRepository) GetByUserID(c context.Context, userID string) (*[]*domain.Like, error) {
	filter := bson.M{"user_id": userID}

	cursor, err := l.database.Collection(l.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	// Iterate through the cursor and decode each document into a Like struct
	var likes []*domain.Like
	for cursor.Next(c) {
		var like domain.Like
		if err := cursor.Decode(&like); err != nil {
			cursor.Close(c)
			return nil, err
		}
		likes = append(likes, &like)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &likes, nil
}


