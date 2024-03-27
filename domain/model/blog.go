package model

type BlogCreate struct{
	Title       string    `json:"title"`
    Content     string    `json:"content"`
	Tags 	  []string  `json:"tags"`
}


type BlogUpdate struct {
	Title       string    `json:"title"`
    Content     string    `json:"content"`
}
type BlogInfo struct{
	BlogID  string `json:"blog_id"`
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	CreatedAt   string `json:"createtimestamp"`
	UpdatedAt    string `json:"updatetimestamp"`
	Tags []string `json:"tags"`
}

type CommentCreate struct{
	Content   string `json:"content"`
}
type RatingCreate struct {
	Rating   int    `json:"rating"`

}