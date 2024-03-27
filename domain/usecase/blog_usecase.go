package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)

type BlogUseCase interface {

	CreateBlog(blog *model.BlogCreate, currUser *model.AuthenticatedUser) (*domain.Blog, error)
	GetBlogs(currUser *model.AuthenticatedUser) (*[]*domain.Blog, error)
	GetBlogByID(blogID string, currUser *model.AuthenticatedUser) (*domain.Blog, error)
	UpdateBlogByID(blogID string, blog *model.BlogCreate, currUser *model.AuthenticatedUser) (*domain.Blog, error)
	DeleteBlogByID(blogID string, currUser *model.AuthenticatedUser) (*domain.Blog, error)

	GetLikesByBlogID(blogID string, currUser *model.AuthenticatedUser) (*[]*domain.Like, error)
	LikeBlogByID(blogID  string, currUser *model.AuthenticatedUser) (*domain.Like, error)
	UnlikeBlogByID(likeID string, currUser *model.AuthenticatedUser) (*domain.Like, error)

	GetCommentsByBlogID(blogID string, currUser *model.AuthenticatedUser) (*[]*domain.Comment, error)
	CreateCommentByBlogID(blogID string, comment *model.CommentCreate, currUser *model.AuthenticatedUser) (*domain.Comment, error)
	UpdateCommentByBlogID(commentID string, comment *model.CommentCreate, currUser *model.AuthenticatedUser) (*domain.Comment, error)
	DeleteCommentByBlogID(commentID string, currUser *model.AuthenticatedUser) (*domain.Comment, error)

	GetSharesByBlogID(blogID string, currUser *model.AuthenticatedUser) (*[]*domain.Share, error)
	ShareBlogByID(blogID string, currUser *model.AuthenticatedUser) (*domain.Share, error)
	UnshareBlogByID(shareID string, currUser *model.AuthenticatedUser) (*domain.Share, error)

	GetRatingsByBlogID(blogId string, currUser *model.AuthenticatedUser) (*[]*domain.BlogRating, error)
	RateBlogByID(blogID string, rating *model.RatingCreate, currUser *model.AuthenticatedUser) (*domain.BlogRating, error)
	UpdateRatingByBlogID(ratingID string, rating *model.RatingCreate, currUser *model.AuthenticatedUser) (*domain.BlogRating, error)
	DeleteRatingByBlogID(ratingID string, currUser *model.AuthenticatedUser) (*domain.BlogRating, error)
}