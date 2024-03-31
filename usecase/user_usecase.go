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
	blogTagRepository domain.BlogTagRepository
}

// FollowUserByID implements usecase.UserUseCase.
func (u *UserUseCase) FollowUserByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Follow, string, error) {
	//TODO : Validation Handling
	follow, err := u.followRepository.Create(u.context, &domain.Follow{
		FollowedID: param.ID,
		FollowerID: currUser.UserID,
	})
	if err != nil {
		return nil, "", err
	}
	return follow, "Followed User Successfully", nil
}

// GetBlogsByID implements usecase.UserUseCase.
func (u *UserUseCase) GetBlogsByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.Blog, string, error) {
	blogs, err := u.blogRepository.GetByUserId(u.context, param.ID)
	if err != nil {
		return nil, "", err
	}
	return blogs, "Blogs Fetched Successfully", nil
}

// GetFollowersByID implements usecase.UserUseCase.
func (u *UserUseCase) GetFollowersByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*model.UserInfo, string, error) {
	//TODO: Aggregation
	follows , err := u.followRepository.GetByFollowedID(u.context, param.ID)
	if err != nil {
		return nil, "", err
	}
	var followers []*model.UserInfo
	for _, follow := range *follows {
		user, err := u.userRepository.GetById(u.context, follow.FollowerID)
		if err != nil {
			continue
		}
		followers = append(followers, &model.UserInfo{
			Username: user.Username,
			Name:     user.Name,
			Bio:      user.Bio,
		})
	}
	return &followers, "Followers Fetched Successfully", nil
}

// GetFollowingsByID implements usecase.UserUseCase.
func (u *UserUseCase) GetFollowingsByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*model.UserInfo, string, error) {
		follows , err := u.followRepository.GetByFollowerID(u.context, param.ID)
	if err != nil {
		return nil, "", err
	}
	var followers []*model.UserInfo
	for _, follow := range *follows {
		user, err := u.userRepository.GetById(u.context, follow.FollowerID)
		if err != nil {
			continue
		}
		followers = append(followers, &model.UserInfo{
			Username: user.Username,
			Name:     user.Name,
			Bio:      user.Bio,
		})
	}
	return &followers, "Follow Feingstched Successfully", nil
}

// GetLikesByID implements usecase.UserUseCase.
func (u *UserUseCase) GetLikesByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.Blog, string, error) {
	like, err := u.likeRepository.GetByUserID(u.context, param.ID)
	if err != nil {
		return nil, "", err
	}
	var blogs []*domain.Blog
	for _, l := range *like {
		blog, err := u.blogRepository.GetByID(u.context, l.BlogID)
		if err != nil {
			continue
		}
		blogs = append(blogs, blog)
	}
	return &blogs, "", nil
}

// GetSharesByID implements usecase.UserUseCase.
func (u *UserUseCase) GetSharesByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.Blog, string, error) {
		shares, err := u.shareRepository.GetByUserID(u.context, param.ID)
	if err != nil {
		return nil, "", err
	}
	var blogs []*domain.Blog
	for _, s := range *shares {
		blog, err := u.blogRepository.GetByID(u.context, s.BlogID)
		if err != nil {
			continue
		}
		blogs = append(blogs, blog)
	}
	return &blogs, "Shares Fetched Successfully", nil
}

// GetUserByID implements usecase.UserUseCase.
func (u *UserUseCase) GetUserByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*model.UserInfo, string, error) {
	user, err := u.userRepository.GetById(u.context, param.ID)
	if err != nil {
		return nil,"", err
	}
	return &model.UserInfo{
		Username: user.Username,
		Name:     user.Name,
		Bio:      user.Bio,
	}, "User Fetched Successfully", nil
}

// GetUsers implements usecase.UserUseCase.
func (u *UserUseCase) GetUsers(currUser *model.AuthenticatedUser, dto any, param *model.SearchParam) (*[]*model.UserInfo, string, error) {
	users, err := u.userRepository.GetAll(u.context, param.Search)
	if err != nil {
		return nil, "", err
	}
	var userInfos []*model.UserInfo
	for _, user := range *users {
		userInfos = append(userInfos, &model.UserInfo{
			Username: user.Username,
			Name:     user.Name,
			Bio:      user.Bio,
		})
	}
	return &userInfos, "Users Fetched Successfully", nil
}

// UnfollowUserByID implements usecase.UserUseCase.
func (u *UserUseCase) UnfollowUserByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Follow, string, error) {
	//TODO : Validation Handling
	follow, err := u.followRepository.Delete(u.context, &domain.Follow{
		FollowedID: param.ID,
		FollowerID: currUser.UserID,
	})
	if err != nil {
		return nil, "", err
	}
	return follow, "Unfollowed User Successfully", nil
}

func NewUserUseCase(context *context.Context, environment *config.Environment, userRepository *domain.UserRepository, followRepository *domain.FollowRepository, blogRepository *domain.BlogRepository, shareRepository *domain.ShareRepository, likeRepository *domain.LikeRepository, blogTagRepository *domain.BlogTagRepository) usecase.UserUseCase {
	return &UserUseCase{
		context:          *context,
		environment:      *environment,
		userRepository:   *userRepository,
		followRepository: *followRepository,
		blogRepository:   *blogRepository,
		shareRepository:  *shareRepository,
		likeRepository:   *likeRepository,
		blogTagRepository: *blogTagRepository,

	}
}


