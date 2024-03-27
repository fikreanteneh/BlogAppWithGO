package domain

import "context"

type Share struct {
	ShareID string `json:"_id" bson:"_id"`
	UserID  string `json:"user_id" bson:"user_id"`
	BlogID  string `json:"blog_id" bson:"blog_id"`
}

type ShareRepository interface {
	Create(c context.Context, share *Share) (*Share, error)
	GetByBlogID(c context.Context, blogID string) (*[]*Share, error)
	GetByUserID(c context.Context, userID string) (*[]*Share, error)
	// TODO: why are we deleting share?

	Delete(c context.Context, shareID string) (*Share, error)
}