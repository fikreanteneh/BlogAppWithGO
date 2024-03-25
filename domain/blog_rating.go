package domain

type BlogRating struct {
	RatingID string `json:"rating_id" bson:"_id"`
	UserID   string `json:"user_id" bson:"user_id"`
	BlogID   string `json:"blog_id"`
	Rating   int    `json:"rating"`
}

type BlogRatingRepository interface {
	InsertRating(rating BlogRating) error
	GetRatingByBlogID(blogID string) (*[]*BlogRating, error)
	GetRatingByUserID(userID string) (*[]*BlogRating, error)
	UpdateRating(rating BlogRating) (*BlogRating, error)
	DeleteRating(ratingID string) (*BlogRating, error)
}