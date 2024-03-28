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
	blogTag.BlogID = primitive.NewObjectID().Hex()

	_, err := b.database.Collection(b.collection).InsertOne(c, *blogTag)
	if err != nil {
		return nil, err
	}

	return blogTag, nil
}

// Delete implements domain.BlogTagRepository.
func (b *BlogTagRepository) Delete(c context.Context, blogTagID string) (*domain.BlogTag, error) {
	filter := bson.M{"_id": blogTagID}
	var blogTag domain.BlogTag
	err := b.database.Collection(b.collection).FindOneAndDelete(c, filter).Decode(&blogTag)
	if err != nil {
		return nil, err
	}

	return &blogTag, nil
}

// GetByBlogID implements domain.BlogTagRepository.
func (b *BlogTagRepository) GetByBlogID(c context.Context, blogID string) (*[]*domain.BlogTag, error) {
	filter := bson.M{"blog_id": blogID}

	// Perform the find operation
	cursor, err := b.database.Collection(b.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	// Iterate through the cursor and decode each document into a User struct
	var blogTags []*domain.BlogTag
	for cursor.Next(c) {
		var blogTag domain.BlogTag
		if err := cursor.Decode(&blogTag); err != nil {
			cursor.Close(c)
			return nil, err
		}
		blogTags = append(blogTags, &blogTag)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &blogTags, nil
}

// GetByTagID implements domain.BlogTagRepository.
func (b *BlogTagRepository) GetByTagID(c context.Context, tagID string) (*[]*domain.BlogTag, error) {
	filter := bson.M{"tag_id": tagID}

	// Perform the find operation
	cursor, err := b.database.Collection(b.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	// Iterate through the cursor and decode each document into a User struct
	var blogTags []*domain.BlogTag
	for cursor.Next(c) {
		var blogTag domain.BlogTag
		if err := cursor.Decode(&blogTag); err != nil {
			cursor.Close(c)
			return nil, err
		}
		blogTags = append(blogTags, &blogTag)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &blogTags, nil
}


