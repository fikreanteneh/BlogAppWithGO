package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
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
	tagRepository domain.TagRepository
	blogTagRepository domain.BlogTagRepository
	commentRepository domain.CommentRepository
}

// CreateBlog implements usecase.BlogUseCase.
func (b *BlogUseCase) CreateBlog(blog *model.BlogCreate, currUser *model.AuthenticatedUser) (*model.BlogInfo, error) {
	//TODO : Authorization Handling
	//TODO : Validation Handling
	createdBlog, err := b.blogRepository.Create(b.context, &domain.Blog{
		UserID:  currUser.UserID,
		Title:   blog.Title,
		Content: blog.Content,

	})
	if err != nil {
		return nil, err
	}
	for _, tag := range blog.Tags {
		b.blogTagRepository.Create(b.context, &domain.BlogTag{
			BlogID: createdBlog.BlogID,
			TagID: tag,
		})
	var tags []string
	for _, tag := range blog.Tags {
		fetchedTag, err := b.tagRepository.GetByID(b.context, tag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, fetchedTag.Name)
		}
	}
	return &model.BlogInfo{
		BlogID: createdBlog.BlogID,
		UserID: createdBlog.UserID,
		Title:  createdBlog.Title,
		Content: createdBlog.Content,
		CreatedAt: createdBlog.CreatedAt.String(),
		UpdatedAt: createdBlog.UpdatedAt.String(),
	}, nil
}

// CreateCommentByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) CreateCommentByBlogID(blogID string, comment *model.CommentCreate, currUser *model.AuthenticatedUser) (*domain.Comment, error) {
	//TODO : Authorization Handling
	//TODO : Validation Handling
	createdComment, err := b.commentRepository.Create(b.context, &domain.Comment{
		UserID:  currUser.UserID,
		BlogID:  blogID,
		Content: comment.Content,

	})
	if err != nil {
		return nil, err
	}
	return createdComment, nil
	
}

// DeleteBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) DeleteBlogByID(blogID string, currUser *model.AuthenticatedUser) (*domain.Blog, error) {
	//TODO : Authorization Handling
	//TODO : Validation Handling
	deletedBlog, err := b.blogRepository.Delete(b.context, blogID)
	if err != nil {
		return nil, err
	}
	return deletedBlog, nil
}

// DeleteCommentByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) DeleteCommentByBlogID(commentID string, currUser *model.AuthenticatedUser) (*domain.Comment, error) {
	//TODO : Authorization Handling
	//TODO : Validation Handling
	deletedComment, err := b.commentRepository.Delete(b.context, commentID)
	if err != nil {
		return nil, err
	}
	return deletedComment, nil
}

// DeleteRatingByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) DeleteRatingByBlogID(ratingID string, currUser *model.AuthenticatedUser) (*domain.BlogRating, error) {
	//TODO : Authorization Handling
	//TODO : Validation Handling
	deletedRating, err := b.ratingRepository.DeleteRating(b.context, ratingID)
	if err != nil {
		return nil, err
	}
	return deletedRating, nil
}

// GetBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) GetBlogByID(blogID string, currUser *model.AuthenticatedUser) (*model.BlogInfo, error) {
	//TODO : Authorization Handling
	//TODO : Validation Handling
	fetchedBlog, err := b.blogRepository.GetByID(b.context, blogID)
	if err != nil {
		return nil, err
	}
	var tags []string
	fetchedBlogTags, err := b.blogTagRepository.GetByBlogID(b.context, blogID)
	if err != nil {
		return nil, err
	}
	for _, fetchedBlogTag := range *fetchedBlogTags {
		fetchedTag, err := b.tagRepository.GetByID(b.context, fetchedBlogTag.TagID)
		if err != nil {
			return nil, err
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
	}, nil
}

// GetBlogs implements usecase.BlogUseCase.
func (b *BlogUseCase) GetBlogs(currUser *model.AuthenticatedUser) (*[]*model.BlogInfo, error) {
	//TODO : Authorization Handling
	//TODO : Validation Handling
	fetchedBlogs, err := b.blogRepository.GetAll(b.context, "")
	if err != nil {
		return nil, err
	}
	var blogInfos []*model.BlogInfo
	for _, fetchedBlog := range *fetchedBlogs {
		var tags []string
		fetchedBlogTags, err := b.blogTagRepository.GetByBlogID(b.context, fetchedBlog.BlogID)
		if err != nil {
			return nil, err
		}
		for _, fetchedBlogTag := range *fetchedBlogTags {
			fetchedTag, err := b.tagRepository.GetByID(b.context, fetchedBlogTag.TagID)
			if err != nil {
				return nil, err
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
	return &blogInfos, nil
}

// GetCommentsByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) GetCommentsByBlogID(blogID string, currUser *model.AuthenticatedUser) (*[]*domain.Comment, error) {
	//TODO : Authorization Handling
	//TODO : Validation Handling
	fetchedComments, err := b.commentRepository.GetByBlogID(b.context, blogID)
	if err != nil {
		return nil, err
	}
	return fetchedComments, nil
}

// GetLikesByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) GetLikesByBlogID(blogID string, currUser *model.AuthenticatedUser) (*[]*model.UserInfo, error) {
	likers, err := b.likeRepository.GetByBlogID(b.context, blogID)
	if err != nil {
		return nil, err
	} 
	var users []*model.UserInfo
	for _, liker := range *likers {
		user, err := b.userRepository.GetById(b.context, liker.UserID)
		if err != nil {
			return nil, err
		}
		users = append(users, &model.UserInfo{
			Username: user.Username,
			Name: user.Name,
			Bio: user.Bio,
			UserId: user.UserID,
		})
	}
	return &users, nil
}

// GetRatingsByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) GetRatingsByBlogID(blogId string, currUser *model.AuthenticatedUser) (*[]*domain.BlogRating, error) {
    ratings, err := b.ratingRepository.GetRatingByBlogID(b.context, blogId)
    if err != nil {
        return nil, err
    }
    return ratings, nil
}

// GetSharesByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) GetSharesByBlogID(blogID string, currUser *model.AuthenticatedUser) (*[]*model.BlogInfo, error) {
    shares, err := b.shareRepository.GetByBlogID(b.context, blogID)
    if err != nil {
        return nil, err
    }
    var blogInfos []*model.BlogInfo
    for _, share := range *shares {
        blog, err := b.blogRepository.GetByID(b.context, share.BlogID)
        if err != nil {
            return nil, err
        }
        blogInfos = append(blogInfos, &model.BlogInfo{
			BlogID: blog.BlogID,
			Content: blog.Content,
			Title: blog.Title,
		})
    }
    return &blogInfos, nil
}
// LikeBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) LikeBlogByID(blogID string, currUser *model.AuthenticatedUser) (*domain.Like, error) {
    likedBlog, err := b.likeRepository.Create(b.context, &domain.Like{
        UserID: currUser.UserID,
        BlogID: blogID,
    })
    if err != nil {
        return nil, err
    }
    return likedBlog, nil
}

// RateBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) RateBlogByID(blogID string, rating *model.RatingCreate, currUser *model.AuthenticatedUser) (*domain.BlogRating, error) {
    ratedBlog, err := b.ratingRepository.InsertRating(b.context, &domain.BlogRating{
        UserID: currUser.UserID,
        BlogID: blogID,
        Rating: rating.Rating,
    })
    if err != nil {
        return nil, err
    }
    return ratedBlog, nil
}

// ShareBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) ShareBlogByID(blogID string, currUser *model.AuthenticatedUser) (*domain.Share, error) {
    sharedBlog, err := b.shareRepository.Create(b.context, &domain.Share{
        UserID: currUser.UserID,
        BlogID: blogID,
    })
    if err != nil {
        return nil, err
    }
    return sharedBlog, nil
}

// UnlikeBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) UnlikeBlogByID(likeID string, currUser *model.AuthenticatedUser) (*domain.Like, error) {
    unlikedBlog, err := b.likeRepository.Delete(b.context, likeID)
    if err != nil {
        return nil, err
    }
    return unlikedBlog, nil
}

// UnshareBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) UnshareBlogByID(shareID string, currUser *model.AuthenticatedUser) (*domain.Share, error) {
    unsharedBlog, err := b.shareRepository.Delete(b.context, shareID)
    if err != nil {
        return nil, err
    }
    return unsharedBlog, nil
}

// UpdateBlogByID implements usecase.BlogUseCase.
func (b *BlogUseCase) UpdateBlogByID(blogID string, blog *model.BlogUpdate, currUser *model.AuthenticatedUser) (*model.BlogInfo, error) {
    _, err := b.blogRepository.Update(b.context, &domain.Blog{
		BlogID: blogID,
        UserID:  currUser.UserID,
        Title:   blog.Title,
        Content: blog.Content,
    })
    if err != nil {
        return nil, err
    }
    return &model.BlogInfo{
		BlogID: blogID,
		Content: blog.Content,
		Title: blog.Title,
	}, nil
}
// UpdateCommentByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) UpdateCommentByBlogID(commentID string, comment *model.CommentCreate, currUser *model.AuthenticatedUser) (*domain.Comment, error) {
    updatedComment, err := b.commentRepository.Update(b.context, &domain.Comment{
		CommentID: commentID,
        UserID:  currUser.UserID,
        Content: comment.Content,
    })
    if err != nil {
        return nil, err
    }
    return updatedComment, nil
}

// UpdateRatingByBlogID implements usecase.BlogUseCase.
func (b *BlogUseCase) UpdateRatingByBlogID(ratingID string, rating *model.RatingCreate, currUser *model.AuthenticatedUser) (*domain.BlogRating, error) {
    updatedRating, err := b.ratingRepository.UpdateRating(b.context, &domain.BlogRating{
		RatingID: ratingID,
        UserID: currUser.UserID,
        Rating: rating.Rating,
    })
    if err != nil {
        return nil, err
    }
    return updatedRating, nil
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
