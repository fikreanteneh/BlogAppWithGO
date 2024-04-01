package domain

import (
	"context"
	"time"
)

type BlogRating struct {
	RatingID string `json:"rating_id" bson:"_id"`
	UserID   string `json:"user_id" bson:"user_id"`
	BlogID   string `json:"blog_id" bson:"blog_id"`
	Rating   int    `json:"rating" bson:"rating"`
	CreatedAt   time.Time `json:"createtimestamp" bson:"createtimestamp"`
	UpdatedAt    time.Time `json:"updatetimestamp" bson:"updatetimestamp"`
}

type BlogRatingRepository interface {
	InsertRating(c context.Context, rating *BlogRating) (*BlogRating, error)
	//create a blog rating and insert it to the database
	GetRatingByBlogID(c context.Context, blogID string) (*[]*BlogRating, error)
	GetRatingByUserID(c context.Context, userID string) (*[]*BlogRating, error)
	UpdateRating(c context.Context, rating *BlogRating) (*BlogRating, error)
	DeleteRating(c context.Context, ratingID string) (*BlogRating, error)
	GetRatingByID(c context.Context, ratingID string) (*BlogRating, error)
	RatingDeleteByBlogID(c context.Context, blogID string)(error)
}