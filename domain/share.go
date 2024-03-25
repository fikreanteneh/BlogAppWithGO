package domain

type Share struct {
	ShareID string `json:"_id" bson:"_id"`
	UserID  string `json:"user_id" bson:"user_id"`
	BlogID  string `json:"blog_id" bson:"blog_id"`
}


type ShareRepository interface {
	Create(share Share) (*Share, error)
	GetByBlogID(blogID string) (*[]*Share, error)
	GetByUserID(userID string) (*[]*Share, error)
	Delete(shareID string) (*Share, error)
}