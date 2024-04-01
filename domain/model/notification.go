package model

import "time"

type NotificationMessage struct {
	NotificationID string `json:"notification_id" bson:"_id"`
	UserID         string `json:"user_id" bson:"user_id"`
	Content        string `json:"content" bson:"content"`
	CreatedAt      time.Time `json:"createtimestamp" bson:"createtimestamp"`
}