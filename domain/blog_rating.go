package domain

import "context"

type BlogRating struct {
	RatingID string `json:"rating_id" bson:"_id"`
	UserID   string `json:"user_id" bson:"user_id"`
	BlogID   string `json:"blog_id" bson:"blog_id"`
	Rating   int    `json:"rating" bson:"rating"`
}

type BlogRatingRepository interface {
	InsertRating(c context.Context, rating *BlogRating) (*BlogRating, error)
	GetRatingByBlogID(c context.Context, blogID string) (*[]*BlogRating, error)
	GetRatingByUserID(c context.Context, userID string) (*[]*BlogRating, error)
	UpdateRating(c context.Context, rating *BlogRating) (*BlogRating, error)
	DeleteRating(c context.Context, ratingID string) (*BlogRating, error)
}