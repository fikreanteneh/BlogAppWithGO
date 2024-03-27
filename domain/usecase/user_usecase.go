package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)
type UserUseCase interface {
    GetUsers() (*[]*model.UserInfo, error)
    GetUserByID(id string) (*model.UserInfo, error)
    GetFollowersByID(id string) (*[]*model.UserInfo, error)
    FollowUserByID(id string) error
    UnfollowUserByID(id string) error
    GetFollowingsByID(id string) (*[]*model.UserInfo, error)
    GetBlogsByID(id string) (*[]*domain.Blog, error)
    GetSharesByID(id string) (*[]*domain.Share, error)
    GetLikesByID(id string) (*[]*domain.Like, error)
}