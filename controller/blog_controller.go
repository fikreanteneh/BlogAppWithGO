package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"BlogApp/middleware"
	"BlogApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogController struct{
	environment *config.Environment
	BlogUseCase usecase.BlogUseCase
}


func NewBlogController(environment *config.Environment, blogUseCase *usecase.BlogUseCase) *BlogController {
	return &BlogController{
		environment: environment,
		BlogUseCase: *blogUseCase,
	}

}

func (b *BlogController) CreateBlog(c *gin.Context){

	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var blogData *model.BlogCreate

    if err := c.BindJSON(&blogData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// TODO: "Blog validation"

	blog, err := b.BlogUseCase.CreateBlog(blogData, currUser)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog"})
        return
    }

	middleware.SuccessResponseHandler(c, http.StatusCreated, "Blog created successfully", blog)
}

func (b *BlogController)GetByBlogID(c *gin.Context){
	blogID := c.Param("id")
    currUser, _ := utils.CheckUser(c)
    blog, err := b.BlogUseCase.GetBlogByID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get blog by ID"})
        return
    }
    
    middleware.SuccessResponseHandler(c, http.StatusOK, "Blog retrieved successfully", blog)
}

func (b *BlogController)GetAllBlogs(c *gin.Context){
    currUser, _ := utils.CheckUser(c)
	blogs, err := b.BlogUseCase.GetBlogs(currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all blogs"})
        return
    }
    
    middleware.SuccessResponseHandler(c, http.StatusOK, "Blogs retrieved successfully", blogs)
}

func (b *BlogController)UpdateBlog(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    var updatedData model.BlogUpdate
    if err := c.BindJSON(&updatedData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	updatedBlog, err := b.BlogUseCase.UpdateBlogByID(blogID, &updatedData, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
        return
    }
    
    middleware.SuccessResponseHandler(c, http.StatusOK, "Blog updated successfully", updatedBlog)
}

func (b BlogController)DeleteBlog(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")

	deleted, err := b.BlogUseCase.DeleteBlogByID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusOK, "Blog deleted successfully", deleted)
}

func (b *BlogController)GetBlogLikes(c *gin.Context){
    currUser, _ := utils.CheckUser(c)
	blogID := c.Param("id")
    
    likes, err := b.BlogUseCase.GetLikesByBlogID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog likes"})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusOK, "Blog likes retrieved successfully", likes)
}

func (b *BlogController) LikeBlog(c *gin.Context){
	blogID := c.Param("id")
    currUser, _ := utils.CheckUser(c)
    
	// TODO : Does the user needs to be authenticated to like/unlike blogs
    like, err := b.BlogUseCase.LikeBlogByID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like blog"})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusCreated, "Blog liked successfully", like)

}

func (b *BlogController) UnlikeBlog(c *gin.Context)  {
	blogID := c.Param("id")
    currUser, _ := utils.CheckUser(c)
    
    like, err := b.BlogUseCase.UnlikeBlogByID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike blog"})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusOK, "Blog unliked successfully", like)
}

func (b *BlogController) GetBlogComments(c *gin.Context) {
	blogID := c.Param("id")
    currUser, _ := utils.CheckUser(c)
    
    comments, err := b.BlogUseCase.GetCommentsByBlogID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog comments"})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusOK, "Blog comments retrieved successfully", comments)
}

func (b *BlogController) UpdateBlogComments(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    
    var data model.CommentCreate
    if err := c.BindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    updatedComments, err := b.BlogUseCase.UpdateCommentByBlogID(blogID, &data, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog comments"})
        return
    }
    
    middleware.SuccessResponseHandler(c, http.StatusAccepted, "Blog comments updated successfully", updatedComments)
}

func (b *BlogController) DeleteBlogComments(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

    commentID := c.Param("comment_id")
    
    comment, err := b.BlogUseCase.DeleteCommentByBlogID(commentID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog comment"})
        return
    }
    
    middleware.SuccessResponseHandler(c, http.StatusOK, "Blog comment deleted successfully", comment)
}

func (b *BlogController) CommentOnBlog(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    
    var comment model.CommentCreate
    if err := c.BindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    addedComment, err := b.BlogUseCase.CreateCommentByBlogID(blogID, &comment, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment to blog"})
        return
    }
    
    middleware.SuccessResponseHandler(c, http.StatusCreated, "Comment added to blog successfully", addedComment)
}

func (b *BlogController) GetBlogShares(c *gin.Context)  {
	blogID := c.Param("id")
    currUser, _ := utils.CheckUser(c)
    
    shares, err := b.BlogUseCase.GetSharesByBlogID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog shares"})
        return
    }

    middleware.SuccessResponseHandler(c, http.StatusOK, "Blog shares retrieved successfully", shares)
}

func (b *BlogController) ShareBlog(c *gin.Context){
	blogID := c.Param("id")
    currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    share, err := b.BlogUseCase.ShareBlogByID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to share blog"})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusOK, "Blog shared successfully", share)
}

func (b *BlogController) UnshareBlog(c *gin.Context){
	shareID := c.Param("share_id")
    currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    share, err := b.BlogUseCase.UnshareBlogByID(shareID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unshare blog"})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusOK, "Blog unshared successfully", share)

}

func (b *BlogController) GetBlogRatings(c *gin.Context){
	blogID := c.Param("id")
    currUser, _ := utils.CheckUser(c)
    
    ratings, err := b.BlogUseCase.GetRatingsByBlogID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog ratings"})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusOK, "Blog ratings retrieved successfully", ratings)
}

func (b *BlogController) RateBlog(c *gin.Context){
	blogID := c.Param("id")
    var rating model.RatingCreate
    if err := c.BindJSON(&rating); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    rate, err := b.BlogUseCase.RateBlogByID(blogID, &rating, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rate blog"})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusCreated, "Blog rated successfully", rate)

}

func (b *BlogController) UpdateBlogRating(c *gin.Context){
	ratingID := c.Param("rating_id")
    currUser, err := utils.CheckUser(c)
    var rate model.RatingCreate
    if err := c.BindJSON(&rate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newRate, err := b.BlogUseCase.UpdateRatingByBlogID(ratingID, &rate, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog rating"})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusOK, "Blog rating updated successfully", newRate)
}