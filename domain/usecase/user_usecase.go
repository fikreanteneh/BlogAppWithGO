package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)
type UserUseCase interface {
    GetUsers(param string) (*[]*model.UserInfo, error)
    GetUserByID(id string) (*model.UserInfo, error)
    GetFollowersByID(id string) (*[]*model.UserInfo, error)
    FollowUserByID(id string, currUser *model.AuthenticatedUser) (*domain.Follow, error)
    UnfollowUserByID(id string, currUser *model.AuthenticatedUser) (*domain.Follow, error)
    GetFollowingsByID(id string) (*[]*model.UserInfo, error) 
    GetBlogsByID(id string) (*[]*domain.Blog, error)
    GetSharesByID(id string) (*[]*domain.Blog, error)
    GetLikesByID(id string) (*[]*domain.Blog, error)
}