package domain

type BlogTag struct {
    BlogTagID string `json:"blog_tag_id" bson:"_id"` 
    BlogID    string `json:"blog_id" bson:"blog_id"`
    TagID     string `json:"tag_id" bson:"tag_id"`
}


type BlogTagRepository interface {
	Create(blogTag BlogTag) (*BlogTag, error)
	GetByBlogID(blogID string) (*[]*BlogTag, error)
	GetByTagID(tagID string) (*[]*BlogTag, error)
	Delete(blogTagID string) (*BlogTag, error)
}

