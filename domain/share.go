package domain

import (
	"context"
	"time"
)

type Share struct {
	ShareID string `json:"_id" bson:"_id"`
	UserID  string `json:"user_id" bson:"user_id"`
	BlogID  string `json:"blog_id" bson:"blog_id"`
	CreatedAt   time.Time `json:"createtimestamp" bson:"createtimestamp"`

}

type ShareRepository interface {
	Create(c context.Context, share *Share) (*Share, error)
	GetByBlogID(c context.Context, blogID string) (*[]*Share, error)
	GetByUserID(c context.Context, userID string) (*[]*Share, error)
	Delete(c context.Context, shareID string) (*Share, error)
	DeleteShareByBlogID(c context.Context, blogID string) error
}