package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
)

type LikeUseCase struct {
	context        context.Context
	environment    config.Environment
	likeRepository domain.LikeRepository
	userRepository domain.UserRepository
}

// GetLikesByBlogID implements usecase.LikeUseCase.
func (l *LikeUseCase) GetLikesByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*model.UserInfo, string, error) {
	likers, err := l.likeRepository.GetByBlogID(l.context, param.ID)
	if err != nil {
		return nil,"Likes Fetching Failed", err
	}
	var users []*model.UserInfo
	for _, liker := range *likers {
		user, err := l.userRepository.GetById(l.context, liker.UserID)
		if err != nil {
			return nil, "Likes Fetching Failed" ,err
		}
		users = append(users, &model.UserInfo{
			Username: user.Username,
			Name: user.Name,
			Bio: user.Bio,
			UserId: user.UserID,
		})
	}
	return &users, "Likes Fetched Successfully", nil
}

// LikeBlogByID implements usecase.LikeUseCase.
func (l *LikeUseCase) LikeBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Like, string, error) {
	    likedBlog, err := l.likeRepository.Create(l.context, &domain.Like{
        UserID: currUser.UserID,
        BlogID: param.ID,
    })
    if err != nil {
        return nil, "Liking Failed", err
    }
    return likedBlog,"Liked a Blog", nil
}

// UnlikeBlogByID implements usecase.LikeUseCase.
func (l *LikeUseCase) UnlikeBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Like, string, error) {

	unlikedBlog, err := l.likeRepository.Delete(l.context, param.ID)

    if err != nil {
        return nil, "Unliking was Unseccessfull",err
    }
    return unlikedBlog, "Unliked a blog successfull",nil
}

func NewLikeUseCase(context *context.Context, environment *config.Environment, likeRepository *domain.LikeRepository, userRepository *domain.UserRepository) usecase.LikeUseCase {
	return &LikeUseCase{
		context:        *context,
		environment:    *environment,
		likeRepository: *likeRepository,
		userRepository: *userRepository,
	}
}
