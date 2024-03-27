package domain

import "context"

type BlogTag struct {
	BlogTagID string `json:"blog_tag_id" bson:"_id"`
	BlogID    string `json:"blog_id" bson:"blog_id"`
	TagID     string `json:"tag_id" bson:"tag_id"`
}

type BlogTagRepository interface {
	Create(c context.Context, blogTag *BlogTag) (*BlogTag, error)
	GetByBlogID(c context.Context, blogID string) (*[]*BlogTag, error)
	GetByTagID(c context.Context, tagID string) (*[]*BlogTag, error)
	Delete(c context.Context, blogTagID string) (*BlogTag, error)
}
