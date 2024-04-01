package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
)

type ShareUseCase struct {
	context         context.Context
	environment     config.Environment
	shareRepository domain.ShareRepository
	blogRepository domain.BlogRepository
}

// GetSharesByBlogID implements usecase.ShareUseCase.
func (s *ShareUseCase) GetSharesByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*model.BlogInfo, string, error) {
	shares, err := s.shareRepository.GetByBlogID(s.context, param.ID)
    if err != nil {
        return nil, "Share Fetching Failed", err
    }
    var blogInfos []*model.BlogInfo
    for _, share := range *shares {
        blog, err := s.blogRepository.GetByID(s.context, share.BlogID)
        if err != nil {
            return nil, "Share Fetching Failed", err
        }
        blogInfos = append(blogInfos, &model.BlogInfo{
			BlogID: blog.BlogID,
			Content: blog.Content,
			Title: blog.Title,
		})
    }
    return &blogInfos,"Share Fetched Successfully", nil
}

// ShareBlogByID implements usecase.ShareUseCase.
func (s *ShareUseCase) ShareBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Share, string, error) {
	    sharedBlog, err := s.shareRepository.Create(s.context, &domain.Share{
        UserID: currUser.UserID,
        BlogID: param.ID,
    })
    if err != nil {
        return nil,"Sharing Blog was unsuccessful", err
    }
    return sharedBlog,"Shared a Blog Successfully", nil
}

// UnshareBlogByID implements usecase.ShareUseCase.
func (s *ShareUseCase) UnshareBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Share, string, error) {
	unsharedBlog, err := s.shareRepository.Delete(s.context, param.ID)
    if err != nil {
        return nil,"Unsharing Failed", err
    }
    return unsharedBlog,"Sharing Successfull", nil
}

func NewShareUseCase(context *context.Context, environment *config.Environment, shareRepository *domain.ShareRepository, blogRepository *domain.BlogRepository) usecase.ShareUseCase {
	return &ShareUseCase{
		context:         *context,
		environment:     *environment,
		shareRepository: *shareRepository,
		blogRepository: *blogRepository,
	}
}
