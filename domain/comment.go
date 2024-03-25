package domain

type Comment struct {
	CommentID string `json:"comment_id" bson:"_id"`
	UserID    string `json:"user_id" bson:"user_id"`
	BlogID    string `json:"blog_id" bson:"blog_id"`
	Content   string `json:"content" bson:"content"`
}


type CommentRepository interface {
	GetByBlogID(blogID string) (*[]*Comment, error)
	Create(comment Comment) (*Comment, error)
	Delete(commentID string) (*Comment, error)
	Update(comment Comment) (*Comment, error)
}