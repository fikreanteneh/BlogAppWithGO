package domain

import "context"

type Tag struct {
	TagID string `json:"_id" bson:"_id"`
	Name  string `json:"name" bson:"_id"`
}

type TagRepository interface {
	Create(c context.Context, tag *Tag) (*Tag, error)
	GetAll(c context.Context, param string) (*[]*Tag, error)
	GetByID(c context.Context, tagID string) (*Tag, error)
	GetByName(c context.Context, name string) (*Tag, error)
	Update(c context.Context, tag *Tag) (*Tag, error)
	Delete(c context.Context, tagID string) (*Tag, error)
	
}