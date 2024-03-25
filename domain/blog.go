package domain

type Blog struct {
    BlogID  string `json:"blog_id" bson:"blog_id"`
    UserID  string `json:"user_id" bson:"user_id"`
    Title   string `json:"title" bson:"title"`
    Content string `json:"content" bson:"content"`
}


type BlogRepository interface {
	GetAll() (*[]*Blog, error)
	GetByID(id string) (*Blog, error)
	GetByUserId(id string) (*[]*Blog, error)
	Create(b Blog) (*Blog,error)
	Update(b Blog) ( *Blog,error)
	Delete(id string) (*Blog, error)
}