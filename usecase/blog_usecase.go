package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
	"errors"
	"time"
)

type BlogUseCase struct {
	context           context.Context
	environment       config.Environment
	blogRepository    domain.BlogRepository
	userRepository    domain.UserRepository
	shareRepository   domain.ShareRepository
	likeRepository    domain.LikeRepository
	ratingRepository  domain.BlogRatingRepository
	tagRepository     domain.TagRepository
	blogTagRepository domain.BlogTagRepository
	commentRepository domain.CommentRepository
}

// CreateBlog implements usecase.BlogUseCase.
func (b *BlogUseCase) CreateBlog(currUser *model.AuthenticatedUser, dto *model.BlogCreate, param any) (*model.BlogInfo, string, error) {
	//TODO : Validation Handling
	createdBlog, err := b.blogRepository.Create(b.context, &domain.Blog{
		UserID:  currUser.UserID,
		Title:   dto.Title,
		Content: dto.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

	})
	if err != nil {
		return nil, "Something Went Wrong", err
	}
	for _, tag := range dto.Tags {
		b.blogTagRepository.Create(b.context, &domain.BlogTag{
			BlogID: createdBlog.BlogID,
			TagID: tag,
		})

	}
	var tags []string
	for _, tag := range dto.Tags {
		fetchedTag, err := b.tagRepository.GetByID(b.context, tag)
		if err != nil {
			return nil, "Tag Does not Exist", err
		}
		tags = append(tags, fetchedTag.Name)

	}
	return &model.BlogInfo{
		BlogID: createdBlog.BlogID,
		UserID: createdBlog.UserID,
		Title:  createdBlog.Title,
		Content: createdBlog.Content,
		CreatedAt: createdBlog.CreatedAt.String(),
		UpdatedAt: createdBlog.UpdatedAt.String(),
		Tags: tags,
	}, "Blog Created",nil
}

// DeleteBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) DeleteBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Blog, string, error) {
	blog, err := b.blogRepository.GetByID(b.context, param.ID)
	if err != nil {
		return nil, "Blog Not Found", err
	}
	if currUser.Role != "ADMIN" || currUser.UserID != blog.UserID {
		return nil, "Unauthorized", errors.New("Unauthorized")
	}
	deletedBlog, err := b.blogRepository.Delete(b.context, param.ID)
	if err != nil {
		return nil,"Blog Deletion UNsuccessful", err
	}
	return deletedBlog, "Blog Deleted Successfully", nil
}

// GetBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) GetBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*model.BlogInfo, string, error) {
	fetchedBlog, err := b.blogRepository.GetByID(b.context, param.ID)
	if err != nil {
		return nil, "Blog Not Found", err
	}
	var tags []string
	fetchedBlogTags, err := b.blogTagRepository.GetByBlogID(b.context, param.ID)
	if err != nil {
		return nil, "Tags Not Found", err
	}
	for _, fetchedBlogTag := range *fetchedBlogTags {
		fetchedTag, err := b.tagRepository.GetByID(b.context, fetchedBlogTag.TagID)
		if err != nil {
			return nil, "Tag Not Found", err
		}
		tags = append(tags, fetchedTag.Name)
	}

	return &model.BlogInfo{
		BlogID: fetchedBlog.BlogID,
		UserID: fetchedBlog.UserID,
		Title:  fetchedBlog.Title,
		Content: fetchedBlog.Content,
		CreatedAt: fetchedBlog.CreatedAt.String(),
		UpdatedAt: fetchedBlog.UpdatedAt.String(),
		Tags: tags,
	}, "Blog Found", nil
}

// GetBlogs implements usecase.BlogUseCase.
func (b *BlogUseCase) GetBlogs(currUser *model.AuthenticatedUser, dto any, param *model.SearchParam) (*[]*model.BlogInfo, string, error) {
	fetchedBlogs, err := b.blogRepository.GetAll(b.context, param.Search)
	if err != nil {
		return nil, "Blogs Not Found", err
	}
	var blogInfos []*model.BlogInfo
	for _, fetchedBlog := range *fetchedBlogs {
		var tags []string
		fetchedBlogTags, err := b.blogTagRepository.GetByBlogID(b.context, fetchedBlog.BlogID)
		if err != nil {
			return nil, "Tags Not Found", err
		}
		for _, fetchedBlogTag := range *fetchedBlogTags {
			fetchedTag, err := b.tagRepository.GetByID(b.context, fetchedBlogTag.TagID)
			if err != nil {
				return nil, "Tag Not Found", err
			}
			tags = append(tags, fetchedTag.Name)
		}
		blogInfos = append(blogInfos, &model.BlogInfo{
			BlogID: fetchedBlog.BlogID,
			UserID: fetchedBlog.UserID,
			Title:  fetchedBlog.Title,
			Content: fetchedBlog.Content,
			CreatedAt: fetchedBlog.CreatedAt.String(),
			UpdatedAt: fetchedBlog.UpdatedAt.String(),
			Tags: tags,
		})
	}
	return &blogInfos, "Blogs Found", nil
}

// UpdateBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) UpdateBlogByID(currUser *model.AuthenticatedUser, dto *model.BlogUpdate, param *model.IdParam) (*model.BlogInfo, string, error) {
	//TODO : Validation Handling
	blog, err := b.blogRepository.GetByID(b.context, param.ID)
	if err != nil {
		return nil, "Blog Not Found", err
	}
	if blog.UserID != currUser.UserID {
		return nil, "Unauthorized", nil
	}
	_, err = b.blogRepository.Update(b.context, &domain.Blog{
		BlogID: param.ID,
		Title:   dto.Title,
		Content: dto.Content,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, "Blog Update Unsuccessful", err
	}
	return &model.BlogInfo{
		BlogID: param.ID,
		Content: dto.Content,
		Title: dto.Title,
	}, "Blog Updated Successfully", nil
}

func NewBlogUseCase(context *context.Context, environment *config.Environment, blogRepository *domain.BlogRepository, userRepository *domain.UserRepository, shareRepository *domain.ShareRepository, likeRepository *domain.LikeRepository, ratingRepository *domain.BlogRatingRepository, tagRepository *domain.TagRepository, blogTagRepository *domain.BlogTagRepository) usecase.BlogUseCase {
	return &BlogUseCase{
		context:          *context,
		environment:      *environment,
		blogRepository:   *blogRepository,
		userRepository:   *userRepository,
		shareRepository:  *shareRepository,
		likeRepository:   *likeRepository,
		ratingRepository: *ratingRepository,
		tagRepository: *tagRepository,
		blogTagRepository: *blogTagRepository,
	}
}














