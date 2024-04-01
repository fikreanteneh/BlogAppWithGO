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

type BlogRatingRepository struct {
	database   *mongo.Database
	collection string
}

// GetRatingByID implements domain.BlogRatingRepository.
func (b *BlogRatingRepository) GetRatingByID(c context.Context, ratingID string) (*domain.BlogRating, error) {
	filter := bson.M{"_id": ratingID}
	var blogRating domain.BlogRating
	err := b.database.Collection(b.collection).FindOne(c, filter).Decode(&blogRating)
	if err != nil {
		return nil, err
	}
	return &blogRating, nil
}

func NewBlogRatingRepository(db *mongo.Database, collection string) domain.BlogRatingRepository {
	return &BlogRatingRepository{
		database:   db,
		collection: collection,
	}
}

// InsertRating implements domain.BlogRatingRepository.
func (b *BlogRatingRepository) InsertRating(c context.Context, rating *domain.BlogRating) (*domain.BlogRating, error) {
	rating.RatingID = primitive.NewObjectID().Hex()
	_, err := b.database.Collection(b.collection).InsertOne(c, *rating)
	if err != nil {
		return nil, err
	}

	return rating, nil
}

// DeleteRating implements domain.BlogRatingRepository.
func (b *BlogRatingRepository) DeleteRating(c context.Context, ratingID string) (*domain.BlogRating, error) {
	filter := bson.M{"_id": ratingID}
	var blogRating domain.BlogRating
	err := b.database.Collection(b.collection).FindOneAndDelete(c, filter).Decode(&blogRating)
	if err != nil {
		return nil, err
	}

	return &blogRating, nil
}

// DeleteByBlogID deletes ratings associated with a specific blog ID.
func (b *BlogRatingRepository) DeleteRatingByBlogID(c context.Context, blogID string) error {
	filter := bson.M{"blog_id": blogID}

	_, err := b.database.Collection(b.collection).DeleteMany(c, filter)
	if err != nil {
		return err
	}
	return nil
}


// GetRatingByBlogID implements domain.BlogRatingRepository.
func (b *BlogRatingRepository) GetRatingByBlogID(c context.Context, blogID string) (*[]*domain.BlogRating, error) {
	filter := bson.M{"blog_id": blogID}

	// Perform the find operation
	cursor, err := b.database.Collection(b.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	// Iterate through the cursor and decode each document into a User struct
	var blogRatings []*domain.BlogRating
	for cursor.Next(c) {
		var blogRating domain.BlogRating
		if err := cursor.Decode(&blogRating); err != nil {
			cursor.Close(c)
			return nil, err
		}
		blogRatings = append(blogRatings, &blogRating)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &blogRatings, nil
}


// add the update rating method
// UpdateRating implements domain.BlogRatingRepository.
func (b *BlogRatingRepository) UpdateRating(c context.Context, rating *domain.BlogRating) (*domain.BlogRating, error) {
	filter := bson.M{"_id": rating.RatingID}
	update := bson.M{"$set": bson.M{"rating": rating.Rating, "updatetimestamp": time.Now()}}

	_, err := b.database.Collection(b.collection).UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}

	return rating, nil
}

// GetRatingByUserID implements domain.BlogRatingRepository.
func (b *BlogRatingRepository) GetRatingByUserID(c context.Context, userID string) (*[]*domain.BlogRating, error) {
	filter := bson.M{"user_id": userID}

	// Perform the find operation
	cursor, err := b.database.Collection(b.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	// Iterate through the cursor and decode each document into a User struct
	var blogRatings []*domain.BlogRating
	for cursor.Next(c) {
		var blogRating domain.BlogRating
		if err := cursor.Decode(&blogRating); err != nil {
			cursor.Close(c)
			return nil, err
		}
		blogRatings = append(blogRatings, &blogRating)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &blogRatings, nil
}
