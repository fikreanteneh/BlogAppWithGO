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

type NotificationRepository struct {
	database   *mongo.Database
	collection string
}

func NewNotificationRepository(db *mongo.Database, collection string) domain.NotificationRepository {
	return &NotificationRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.NotificationRepository.
func (n *NotificationRepository) Create(c context.Context, notification *domain.Notification) (*domain.Notification, error) {
	notification.NotificationID = primitive.NewObjectID().Hex()
	  _, err := n.database.Collection(n.collection).InsertOne(c, *notification)
	  if err != nil{ 
		return nil, err
	  }
	
	  return notification, nil
	
}

// Delete implements domain.NotificationRepository.
func (n *NotificationRepository) Delete(c context.Context, notificationID string) (*domain.Notification, error) {
	filter := bson.M{"_id": notificationID}

	var notification domain.Notification
	err := n.database.Collection(n.collection).FindOneAndDelete(c, filter).Decode(&notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

// GetByUserId implements domain.NotificationRepository.
func (n *NotificationRepository) GetByUserId(c context.Context, userID string) (*[]*domain.Notification, error) {
	filter := bson.M{"user_id": userID}

	cursor, err := n.database.Collection(n.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	// Iterate through the cursor and decode each document into a Like struct
	var notifications []*domain.Notification
	for cursor.Next(c) {
		var notification domain.Notification
		if err := cursor.Decode(&notification); err != nil {
			cursor.Close(c)
			return nil, err
		}
		notifications = append(notifications, &notification)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &notifications, nil
}


