package domain

import (
	"context"
	"time"
)

type Like struct {
	LikeID string `json:"like_id" bson:"_id"`
	UserID string `json:"user_id" bson:"user_id"`
	BlogID string `json:"blog_id" bson:"blog_id"`
	CreatedAt   time.Time `json:"createtimestamp" bson:"createtimestamp"`

}

type LikeRepository interface {
	Create(c context.Context, like *Like) (*Like, error)
	GetByBlogID(c context.Context, blogID string) (*[]*Like, error)
	GetByUserID(c context.Context, userID string) (*[]*Like, error)
	Delete(c context.Context, likeID string) (*Like, error)
	DeleteByBlogId(c context.Context, blogID string) (any, error)
}