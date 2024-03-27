package usecase

import (
	"BlogApp/domain"
)

type BlogUseCase interface {
	GetBlogs() (*[]*domain.Blog, error)
	GetBlogByID(id string) (*domain.Blog, error)
	CreateBlog(blog *domain.Blog) error
	UpdateBlogByID(id string, blog *domain.Blog) error
	DeleteBlogByID(id string) error

	GetLikesByBlogID(id string) (*[]*domain.Like, error)
	LikeBlogByID(id string) error
	UnlikeBlogByID(id string) error

	GetCommentsByBlogID(id string) (*[]*domain.Comment, error)
	CreateCommentByBlogID(id string, comment *domain.Comment) error
	UpdateCommentByBlogID(id string, comment *domain.Comment) error
	DeleteCommentByBlogID(id string) error

	GetSharesByBlogID(id string) (*[]*domain.Share, error)
	ShareBlogByID(id string) error
	UnshareBlogByID(id string) error

	GetRatingsByBlogID(id string) (*[]*domain.BlogRating, error)
	RateBlogByID(id string, rating *domain.BlogRating) error
	UpdateRatingByBlogID(id string, rating *domain.BlogRating) error
	DeleteRatingByBlogID(id string) error
}