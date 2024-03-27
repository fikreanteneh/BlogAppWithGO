package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/model"
)

type BlogUseCase interface {
	GetBlogs() (*[]*domain.Blog, error)
	GetBlogByID(id string) (*domain.Blog, error)
	CreateBlog(blog *domain.Blog, currUser model.AuthenticatedUser) (*domain.Blog, error)
	UpdateBlogByID(id string, blog *domain.Blog) (*domain.Blog, error)
	DeleteBlogByID(id string) (*domain.Blog, error)
	GetLikesByBlogID(id string) (*[]*domain.Like, error)
	LikeBlogByID(id string) (*domain.Blog, error)
	UnlikeBlogByID(id string) (*domain.Blog, error)

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