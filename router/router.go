package router

import (
	"BlogApp/config"
	"BlogApp/controller"
	"BlogApp/domain"
	"BlogApp/middleware"
	"BlogApp/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *config.Environment, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	var userRepository domain.UserRepository
	var blogRepository domain.BlogRepository
	var tagRepository domain.TagRepository
	var notificationRepository domain.NotificationRepository
	var blogRatingRepository domain.BlogRatingRepository
	var blogShareRepository domain.ShareRepository
	var blogCommentRepository domain.CommentRepository
	var blogLikeRepository domain.LikeRepository
	var tagBlogRepository domain.TagRepository
	var followRepository domain.FollowRepository

	userUseCase := usecase.NewAuthUseCase(env, &userRepository)

	var authController = controller.NewAuthController(env, &userUseCase)

	publicRouter := gin.Group("auth")
	publicRouter.POST("/register", authController.Register)
	publicRouter.POST("/login", authController.Login)
	publicRouter.POST("/adminRegister", authController.AdminRegister)

	profileRouter := gin.Group("profile")
	profileRouter.Use(middleware.AuthMiddleware(env.JwtSecret))
	profileRouter.GET("/")
	profileRouter.PUT("/")
	profileRouter.PUT("/updatePassword") 
	profileRouter.PUT("/updateEmail") 
	profileRouter.PUT("/updateUsername") 
	profileRouter.PUT("/updateProfilePicture")
	profileRouter.DELETE("/")

	userRouter := gin.Group("user")
	userRouter.Use(middleware.AuthMiddleware(env.JwtSecret))
	userRouter.GET("/")
	userRouter.GET("/:id")
	userRouter.GET("/:id/followers")
	userRouter.POST("/:id/follow")
	userRouter.DELETE("/:id/follow")
	userRouter.GET("/:id/followings")
	userRouter.GET("/:id/blogs")
	userRouter.GET("/:id/shares")
	userRouter.GET("/:id/likes")

	notificationRouter := gin.Group("notification")
	userRouter.Use(middleware.AuthMiddleware(env.JwtSecret))
	notificationRouter.GET("/")
	notificationRouter.GET("/:id")
	notificationRouter.DELETE("/:id")

	blogRouter := gin.Group("blog")
	blogRouter.Use(middleware.AuthMiddleware(env.JwtSecret))
	blogRouter.GET("/")
	blogRouter.GET("/:id")
	blogRouter.POST("/")
	blogRouter.PUT("/:id")
	blogRouter.DELETE("/:id")

	blogRouter.GET("/:id/likes")
	blogRouter.POST("/:id/likes")
	blogRouter.DELETE("/:id/likes")

	blogRouter.GET("/:id/comments")
	blogRouter.POST("/:id/comments")
	blogRouter.PUT("/:id/comments")
	blogRouter.DELETE("/:id/comments")

	blogRouter.GET("/:id/shares")
	blogRouter.POST("/:id/shares")
	blogRouter.DELETE("/:id/shares")

	blogRouter.GET("/:id/ratings")
	blogRouter.POST("/:id/ratings")
	blogRouter.PUT("/:id/ratings")
	blogRouter.DELETE("/:id/ratings")

	tagRouter := gin.Group("tag")
	tagRouter.Use(middleware.AuthMiddleware(env.JwtSecret))
	tagRouter.GET("/")
	tagRouter.GET("/:id")
	tagRouter.POST("/")
	tagRouter.PUT("/:id")
	tagRouter.DELETE("/:id")
}
