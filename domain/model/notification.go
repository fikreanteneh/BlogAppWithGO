package model

import "time"

type NotificationMessage struct {
	Content        string `json:"content" bson:"content"`
	CreatedAt      time.Time `json:"createtimestamp" bson:"createtimestamp"`
}