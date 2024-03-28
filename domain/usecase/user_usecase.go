package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)
type UserUseCase interface {
    GetUsers(currUser *model.AuthenticatedUser, dto any, param *model.SearchParam) (*[]*model.UserInfo, string, error)
    GetUserByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*model.UserInfo, string, error)
    GetFollowersByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*model.UserInfo, string, error)
    FollowUserByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Follow, string, error)
    UnfollowUserByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Follow, string, error)
    GetFollowingsByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*model.UserInfo, string, error) 
    GetBlogsByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.Blog, string, error)
    GetSharesByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.Blog, string, error)
    GetLikesByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.Blog, string, error)
}