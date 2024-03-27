package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
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

	c.JSON(http.StatusCreated, gin.H{
        "message": "Blog created successfully",
        "blog":    blog,
    })
}

func (b *BlogController)GetByBlogID(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    
    blog, err := b.BlogUseCase.GetBlogByID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get blog by ID"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Blog retrieved successfully",
        "blog":    blog,
    })
}

func (b *BlogController)GetAllBlogs(c *gin.Context){
	
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogs, err := b.BlogUseCase.GetBlogs(currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all blogs"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "All blogs retrieved successfully",
        "blogs":    blogs,
    })
}

func (b *BlogController)UpdateBlog(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    var updatedData model.BlogCreate
    if err := c.BindJSON(&updatedData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	updatedBlog, err := b.BlogUseCase.UpdateBlogByID(blogID, &updatedData, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Blog updated successfully",
        "blog":    updatedBlog,
    })
}

func (b BlogController)DeleteBlog(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")

	_ , err = b.BlogUseCase.DeleteBlogByID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func (b *BlogController)GetBlogLikes(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	
	blogID := c.Param("id")
    
    likes, err := b.BlogUseCase.GetLikesByBlogID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog likes"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Blog likes retrieved successfully",
        "likes":   likes,
    })
}

func (b *BlogController) LikeBlog(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    
	// TODO : Does the user needs to be authenticated to like/unlike blogs
    _, err = b.BlogUseCase.LikeBlogByID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like blog"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Blog liked successfully"})
}

func (b *BlogController) UnlikeBlog(c *gin.Context)  {
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    
    _, err = b.BlogUseCase.UnlikeBlogByID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike blog"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Blog unliked successfully"})
}

func (b *BlogController) GetBlogComments(c *gin.Context) {
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    
    comments, err := b.BlogUseCase.GetCommentsByBlogID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog comments"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message":  "Blog comments retrieved successfully",
        "comments": comments,
    })
}

func (b *BlogController) UpdateBlogComments(c *gin.Context){

	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	commentID := c.Param("id")
    
    var updatedComments *model.CommentCreate
    if err := c.BindJSON(&updatedComments); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    _ , err = b.BlogUseCase.UpdateCommentByBlogID(commentID, updatedComments, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog comments"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Blog comments updated successfully"})
}

func (b *BlogController) DeleteBlogComments(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

    commentID := c.Param("comment_id")
    
    _, err = b.BlogUseCase.DeleteCommentByBlogID(commentID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog comment"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Blog comment deleted successfully"})
}

func (b *BlogController) CommentOnBlog(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    
    var comment *model.CommentCreate
    if err := c.BindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    addedComment, err := b.BlogUseCase.CreateCommentByBlogID(blogID, comment, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment to blog"})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{
        "message": "Comment added to blog successfully",
        "comment": addedComment,
    })
}

func (b *BlogController) GetBlogShares(c *gin.Context)  {
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    
    shares, err := b.BlogUseCase.GetSharesByBlogID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog shares"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Blog shares retrieved successfully",
        "shares": shares,
    })
}

func (b *BlogController) ShareBlog(c *gin.Context){
	blogID := c.Param("id")
    currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    _,err = b.BlogUseCase.ShareBlogByID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to share blog"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog shared successfully"})
}

func (b *BlogController) UnshareBlog(c *gin.Context){
	shareID := c.Param("id")
    currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    _, err = b.BlogUseCase.UnshareBlogByID(shareID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unshare blog"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog unshared successfully"})
}

func (b *BlogController) GetBlogRatings(c *gin.Context){
	currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

	blogID := c.Param("id")
    
    ratings, err := b.BlogUseCase.GetRatingsByBlogID(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog ratings"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Blog ratings retrieved successfully",
        "ratings": ratings,
    })
}

func (b *BlogController) RateBlog(c *gin.Context){
	blogID := c.Param("id")
    currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

	var rating *model.RatingCreate

    _, err = b.BlogUseCase.RateBlogByID(blogID, rating, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rate blog"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog rated successfully"})
}

func (b *BlogController) UpdateBlogRating(c *gin.Context){
	ratingID := c.Param("id")
    currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }
    
	var newRating *model.RatingCreate
    _, err = b.BlogUseCase.UpdateRatingByBlogID(ratingID, newRating, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog rating"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog rating updated successfully"})
}

func (b *BlogController) DeleteBlogRating(c *gin.Context){
	currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

	ratingID := c.Param("id")
    
    _, err = b.BlogUseCase.DeleteRatingByBlogID(ratingID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog ratings"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{ "message": "Blog ratings deleted successfully" })
}