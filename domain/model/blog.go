package model

type BlogCreate struct{
	Title       string    `json:"title"`
    Content     string    `json:"content"`
}

type CommentCreate struct{
	
	Content   string `json:"content"`
}
type RatingCreate struct {

	Rating   int    `json:"rating"`

}