package domain

import "context"

type Blog struct {
	BlogID  string `json:"blog_id" bson:"blog_id"`
	UserID  string `json:"user_id" bson:"user_id"`
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
}

type BlogRepository interface {
	GetAll(c context.Context, param string) (*[]*Blog, error)
	GetByID(c context.Context, blogID string) (*Blog, error)
	GetByUserId(c context.Context, userID string) (*[]*Blog, error)
	Create(c context.Context, blog *Blog) (*Blog, error)
	Update(c context.Context, blog *Blog) (*Blog, error)
	Delete(c context.Context, blogID string) (*Blog, error)
}