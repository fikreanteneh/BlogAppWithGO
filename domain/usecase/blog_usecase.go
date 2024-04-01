package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)

type BlogUseCase interface {

	CreateBlog(currUser *model.AuthenticatedUser,dto *model.BlogCreate, param any) (*model.BlogInfo, string, error)
	GetBlogs(currUser *model.AuthenticatedUser, dto any, param *model.SearchParam) (*[]*model.BlogInfo, string, error)
	GetBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*model.BlogInfo, string, error)
	UpdateBlogByID(currUser *model.AuthenticatedUser, dto *model.BlogUpdate, param *model.IdParam) (*model.BlogInfo, string, error)
	DeleteBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Blog, string, error)
}