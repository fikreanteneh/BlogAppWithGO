
package domain

type Like struct {
	LikeID string `json:"like_id" bson:"_id"`
	UserID string `json:"user_id" bson:"user_id"`
	BlogID string `json:"blog_id" bson:"blog_id"`
}

type LikeRepository interface {
	Create(like Like) (*Like, error)
	GetByBlogID(blogID string) (*[]*Like, error)
	GetByUserID(userID string) (*[]*Like, error)
	Delete(likeID string) (*Like, error)
}