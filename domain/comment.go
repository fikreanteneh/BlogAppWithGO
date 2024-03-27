package domain

import "context"

type Comment struct {
	CommentID string `json:"comment_id" bson:"_id"`
	UserID    string `json:"user_id" bson:"user_id"`
	BlogID    string `json:"blog_id" bson:"blog_id"`
	Content   string `json:"content" bson:"content"`
}

type CommentRepository interface {
	GetByBlogID(c context.Context, blogID string) (*[]*Comment, error)
	Create(c context.Context, comment *Comment) (*Comment, error)
	Delete(c context.Context, commentID string) (*Comment, error)
	Update(c context.Context, comment *Comment) (*Comment, error)
}