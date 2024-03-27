package usecase

import (
	"BlogApp/domain"
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
	CreateCommentByBlogID(id string, comment *domain.Comment) (*domain.Blog, error)
	UpdateCommentByBlogID(id string, comment *domain.Comment) (*domain.Blog, error)
	DeleteCommentByBlogID(id string) (*domain.Blog, error)

	GetSharesByBlogID(id string) (*[]*domain.Share, error)
	ShareBlogByID(id string) (*domain.Blog, error)
	UnshareBlogByID(id string) (*domain.Blog, error)

	GetRatingsByBlogID(id string) (*[]*domain.BlogRating, error)
	RateBlogByID(id string, rating *domain.BlogRating) (*domain.Blog, error)
	UpdateRatingByBlogID(id string, rating *domain.BlogRating) (*domain.Blog, error)
	DeleteRatingByBlogID(id string) (*domain.Blog, error)
}