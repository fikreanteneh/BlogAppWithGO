package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/usecase"
	"context"
)

type BlogUseCase struct {
	context          context.Context
	environment      config.Environment
	blogRepository   domain.BlogRepository
	userRepository   domain.UserRepository
	shareRepository  domain.ShareRepository
	likeRepository   domain.LikeRepository
	ratingRepository domain.BlogRatingRepository
}

// CreateBlog implements usecase.BlogUseCase.
func (b *BlogUseCase) CreateBlog(blog *domain.Blog) (*domain.Blog, error) {
	panic("unimplemented")
}

// CreateCommentByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) CreateCommentByBlogID(id string, comment *domain.Comment) (*domain.Blog, error) {
	panic("unimplemented")
}

// DeleteBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) DeleteBlogByID(id string) (*domain.Blog, error) {
	panic("unimplemented")
}

// DeleteCommentByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) DeleteCommentByBlogID(id string) (*domain.Blog, error) {
	panic("unimplemented")
}

// DeleteRatingByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) DeleteRatingByBlogID(id string) (*domain.Blog, error) {
	panic("unimplemented")
}

// GetBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) GetBlogByID(id string) (*domain.Blog, error) {
	panic("unimplemented")
}

// GetBlogs implements usecase.BlogUseCase.
func (b *BlogUseCase) GetBlogs() (*[]*domain.Blog, error) {
	panic("unimplemented")
}

// GetCommentsByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) GetCommentsByBlogID(id string) (*[]*domain.Comment, error) {
	panic("unimplemented")
}

// GetLikesByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) GetLikesByBlogID(id string) (*[]*domain.Like, error) {
	panic("unimplemented")
}

// GetRatingsByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) GetRatingsByBlogID(id string) (*[]*domain.BlogRating, error) {
	panic("unimplemented")
}

// GetSharesByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) GetSharesByBlogID(id string) (*[]*domain.Share, error) {
	panic("unimplemented")
}

// LikeBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) LikeBlogByID(id string) (*domain.Blog, error) {
	panic("unimplemented")
}

// RateBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) RateBlogByID(id string, rating *domain.BlogRating) (*domain.Blog, error) {
	panic("unimplemented")
}

// ShareBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) ShareBlogByID(id string) (*domain.Blog, error) {
	panic("unimplemented")
}

// UnlikeBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) UnlikeBlogByID(id string) (*domain.Blog, error) {
	panic("unimplemented")
}

// UnshareBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) UnshareBlogByID(id string) (*domain.Blog, error) {
	panic("unimplemented")
}

// UpdateBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) UpdateBlogByID(id string, blog *domain.Blog) (*domain.Blog, error) {
	panic("unimplemented")
}

// UpdateCommentByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) UpdateCommentByBlogID(id string, comment *domain.Comment) (*domain.Blog, error) {
	panic("unimplemented")
}

// UpdateRatingByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) UpdateRatingByBlogID(id string, rating *domain.BlogRating) (*domain.Blog, error) {
	panic("unimplemented")
}

func NewBlogUseCase(environment *config.Environment, blogRepository *domain.BlogRepository, userRepository *domain.UserRepository, shareRepository *domain.ShareRepository, likeRepository *domain.LikeRepository, ratingRepository *domain.BlogRatingRepository) usecase.BlogUseCase {
	return &BlogUseCase{
		environment:      *environment,
		blogRepository:   *blogRepository,
		userRepository:   *userRepository,
		shareRepository:  *shareRepository,
		likeRepository:   *likeRepository,
		ratingRepository: *ratingRepository,
	}
}
