package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
)

type UserUseCase struct {
	context          context.Context
	environment      config.Environment
	userRepository   domain.UserRepository
	followRepository domain.FollowRepository
	blogRepository   domain.BlogRepository
	shareRepository  domain.ShareRepository
	likeRepository   domain.LikeRepository
}

// FollowUserByID implements usecase.UserUseCase.
func (u *UserUseCase) FollowUserByID(id string, currUser *model.AuthenticatedUser) (*domain.Follow, error) {
	//TODO : Validation Handling 
	follow, err := u.followRepository.Create(u.context, &domain.Follow{
		FollowedID: id,
		FollowerID: currUser.UserID,
	})
	if err != nil {
		 return nil, err
	}
	return follow, nil
}

// GetBlogsByID implements usecase.UserUseCase.
func (u *UserUseCase) GetBlogsByID(id string) (*[]*domain.Blog, error) {
	blogs, err := u.blogRepository.GetAll(u.context, id)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

// GetFollowersByID implements usecase.UserUseCase.
func (u *UserUseCase) GetFollowersByID(id string) (*[]*model.UserInfo, error) {
	//TODO: Aggregation
	//TODO: Validation

	follows, err := u.followRepository.GetByFollowerID(u.context, id)
	if err != nil {
		return nil, err
	}
	var followers []*model.UserInfo
	for _, follow := range *follows {
		user, err := u.userRepository.GetById(u.context, follow.FollowedID)
		if err != nil {
			continue
		}
		followers = append(followers, &model.UserInfo{
			Username: user.Username,
			Name:     user.Name,
			Bio:      user.Bio,
		})
	
	}
	return &followers, nil
}

// GetFollowingsByID implements usecase.UserUseCase.
func (u *UserUseCase) GetFollowingsByID(id string) (*[]*model.UserInfo, error) {
	follows, err := u.followRepository.GetByFollowedID(u.context, id)
	if err != nil {
		return nil, err
	}
	var followings []*model.UserInfo
	for _, follow := range *follows {
		user, err := u.userRepository.GetById(u.context, follow.FollowedID)
		if err != nil {
			continue
		}
		followings = append(followings, &model.UserInfo{
			Username: user.Username,
			Name:     user.Name,
			Bio:      user.Bio,
		})
	
	}
	return &followings, nil
}

// GetLikesByID implements usecase.UserUseCase.
func (u *UserUseCase) GetLikesByID(id string) (*[]*domain.Blog, error) {
	like, err := u.likeRepository.GetByUserID(u.context, id)
	if err != nil {
		return nil, err
	}
	var blogs []*domain.Blog
	for _, l := range *like {
		blog, err := u.blogRepository.GetByID(u.context, l.BlogID)
		if err != nil {
			continue
		}
		blogs = append(blogs, blog)
	}
	return &blogs, nil
}

// GetSharesByID implements usecase.UserUseCase.
func (u *UserUseCase) GetSharesByID(id string) (*[]*domain.Blog, error) {
	shares, err := u.shareRepository.GetByUserID(u.context, id)
	if err != nil {
		return nil, err
	}
	var blogs []*domain.Blog
	for _, s := range *shares {
		blog, err := u.blogRepository.GetByID(u.context, s.BlogID)
		if err != nil {
			continue
		}
		blogs = append(blogs, blog)
	}
	return &blogs, nil
}

// GetUserByID implements usecase.UserUseCase.
func (u *UserUseCase) GetUserByID(id string) (*model.UserInfo, error) {
	user, err := u.userRepository.GetById(u.context, id)
	if err != nil {
		return nil, err
	}
	return &model.UserInfo{
		Username: user.Username,
		Name:     user.Name,
		Bio:      user.Bio,
	}, nil
}

// GetUsers implements usecase.UserUseCase.
func (u *UserUseCase) GetUsers(param string) (*[]*model.UserInfo, error) {
	users, err := u.userRepository.GetAll(u.context, param)
	if err != nil {
		return nil, err
	}
	var userInfos []*model.UserInfo
	for _, user := range *users {
		userInfos = append(userInfos, &model.UserInfo{
			Username: user.Username,
			Name:     user.Name,
			Bio:      user.Bio,
		})
	}
	return &userInfos, nil
}

// UnfollowUserByID implements usecase.UserUseCase.
func (u *UserUseCase) UnfollowUserByID(id string, currUser *model.AuthenticatedUser) (*domain.Follow, error) {
	//TODO : Validation Handling
	follow, err := u.followRepository.Delete(u.context, &domain.Follow{
		FollowedID: id,
		FollowerID: currUser.UserID,
	})
	if err != nil {
		return nil, err
	}
	_, err = u.followRepository.Delete(u.context, follow)
	if err != nil {
		return nil, err
	}
	return follow, nil
}

func NewUserUseCase(environment *config.Environment, userRepository *domain.UserRepository, followRepository *domain.FollowRepository, blogRepository *domain.BlogRepository, shareRepository *domain.ShareRepository, likeRepository *domain.LikeRepository) usecase.UserUseCase {
	return &UserUseCase{
		environment:      *environment,
		userRepository:   *userRepository,
		followRepository: *followRepository,
		blogRepository:   *blogRepository,
		shareRepository:  *shareRepository,
		likeRepository:   *likeRepository,
	}
}
