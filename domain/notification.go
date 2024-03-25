package domain

type Notification struct {
	NotificationID string `json:"notification_id" bson:"_id"`
	UserID         string `json:"user_id" bson:"user_id"`
	Content        string `json:"content" bson:"content"`
}


type NotificationRepository interface {
	GetByUserId() (*[]*Notification, error)
	Create(n Notification) (*Notification, error)
	Delete(id string) (*Notification, error)
}