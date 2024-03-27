package repository

import (
	"BlogApp/domain"
	"context"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
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
	
}

// Delete implements domain.NotificationRepository.
func (n *NotificationRepository) Delete(c context.Context, notificationID string) (*domain.Notification, error) {
	
}

// GetByUserId implements domain.NotificationRepository.
func (n *NotificationRepository) GetByUserId(c context.Context, userID string) (*[]*domain.Notification, error) {
	
}


