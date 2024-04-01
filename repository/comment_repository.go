package repository

import (
	"BlogApp/domain"
	"context"
	"fmt"
	"time"

	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetByID implements domain.CommentRepository.
func (m *CommentRepository) GetByID(c context.Context, commentId string) (*domain.Comment, error) {
	filter := bson.M{"_id": commentId}
	var comment domain.Comment
	err := m.database.Collection(m.collection).FindOne(c, filter).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// Create implements domain.CommentRepository.
func (m *CommentRepository) Create(c context.Context, comment *domain.Comment) (*domain.Comment, error) {
	comment.CommentID = primitive.NewObjectID().Hex()

	_, err := m.database.Collection(m.collection).InsertOne(c, *comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// Delete implements domain.CommentRepository.
func (m *CommentRepository) Delete(c context.Context, commentID string) (*domain.Comment, error) {
	filter := bson.M{"_id": commentID}
	var comment domain.Comment
	err := m.database.Collection(m.collection).FindOneAndDelete(c, filter).Decode(&comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// GetByBlogID implements domain.CommentRepository.
func (m *CommentRepository) GetByBlogID(c context.Context, blogID string) (*[]*domain.Comment, error) {
	filter := bson.M{"blog_id": blogID}

	// Perform the find operation
	cursor, err := m.database.Collection(m.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	// Iterate through the cursor and decode each document into a User struct
	var Comments []*domain.Comment
	for cursor.Next(c) {
		var comment domain.Comment
		if err := cursor.Decode(&comment); err != nil {
			cursor.Close(c)
			return nil, err
		}
		Comments = append(Comments, &comment)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &Comments, nil

}

// Update implements domain.CommentRepository.
func (m *CommentRepository) Update(c context.Context, comment *domain.Comment) (*domain.Comment, error) {
	filter := bson.M{"_id": comment.CommentID}
	update := bson.M{
		"$set": bson.M{
			"content":    comment.Content,
			"updated_at": time.Now(),
		},
	}
	_, err := m.database.Collection(m.collection).UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}

	return comment, nil
}


// DeleteByBlogId implements domain.CommentRepository.
func (m *CommentRepository) DeleteByBlogId(c context.Context, blogID string) (any, error) {
	filter := bson.M{"blog_id": blogID}
	_, err := m.database.Collection(m.collection).DeleteMany(c, filter)
	if err != nil {
		fmt.Println("=========", err)
		return nil, err
	}
	print("========= Deleted SUccessfully")
	return nil, nil
}