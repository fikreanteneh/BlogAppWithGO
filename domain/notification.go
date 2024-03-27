package domain

import (
	"context"
	"time"
)

type Notification struct {
	NotificationID string `json:"notification_id" bson:"_id"`
	UserID         string `json:"user_id" bson:"user_id"`
	Content        string `json:"content" bson:"content"`
	CreatedAt   time.Time `json:"createtimestamp" bson:"createtimestamp"`

}

type NotificationRepository interface {
	GetByUserId(c context.Context, userID string) (*[]*Notification, error)
	Create(c context.Context, notification *Notification) (*Notification, error)
	Delete(c context.Context, notificationID string) (*Notification, error)
}