package router

import (
	"BlogApp/config"
	"BlogApp/controller"
	"BlogApp/middleware"
	"BlogApp/repository"
	"BlogApp/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *config.Environment, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	userRepository := repository.NewUserRepository(db, "users")
	blogRepository := repository.NewBlogRepository(db, "blogs")
	// tagRepository := repository.NewTagRepository(db, "tags")
	// notificationRepository := repository.NewNotificationRepository(db, "notifications")
	shareRepository := repository.NewShareRepository(db, "shares")
	// blogTagRepository := repository.NewBlogTagRepository(db, "blogTags")
	// commentRepository := repository.NewCommentRepository(db, "comments")
	followRepository := repository.NewFollowRepository(db, "follows")
	likeRepository := repository.NewLikeRepository(db, "likes")



	authUseCase := usecase.NewAuthUseCase(env, &userRepository )
	profileUsecase := usecase.NewProfileUseCase(env, &userRepository)
	userUseCase := usecase.NewUserUseCase(env, &userRepository, &followRepository, &blogRepository, &shareRepository, &likeRepository)
	// blogUseCase := 
	
	authController := controller.NewAuthController(env, &authUseCase)
	profileController :=  controller.NewProfileController(env, &profileUsecase)
	userController := controller.NewUserController(env, &userUseCase)
	// blogController := controller.NewBlogController(env, &blogUseCase)

	publicRouter := gin.Group("auth")
	publicRouter.POST("/register", authController.Register)
	publicRouter.POST("/login", authController.Login)
	publicRouter.POST("/adminRegister", authController.AdminRegister)

	profileRouter := gin.Group("profile")
	profileRouter.Use(middleware.AuthMiddleware(env.JwtSecret))
	profileRouter.GET("/", profileController.GetProfile)
	profileRouter.PUT("/", profileController.UpdateProfile)
	profileRouter.DELETE("/", profileController.DeleteProfile)
	profileRouter.PUT("/updatePassword", profileController.UpdatePassword) 
	profileRouter.PUT("/updateEmail", profileController.UpdateEmail) 
	profileRouter.PUT("/updateUsername", profileController.UpdateUsername) 
	//TODO: Implement UpdateProfilePicture
	profileRouter.PUT("/updateProfilePicture", profileController.UpdateProfilePicture)

	userRouter := gin.Group("user")
	userRouter.Use(middleware.AuthMiddleware(env.JwtSecret))
	userRouter.GET("/", userController.GetUsers)
	userRouter.GET("/:id", userController.GetUserByID)
	userRouter.GET("/:id/followers", userController.GetFollowersByID)
	userRouter.POST("/:id/follow", userController.FollowUserByID)
	userRouter.DELETE("/:id/follow", userController.UnfollowUserByID)
	userRouter.GET("/:id/followings", userController.GetFollowingsByID)
	userRouter.GET("/:id/blogs", userController.GetBlogsByID)
	userRouter.GET("/:id/shares", userController.GetSharesByID)
	userRouter.GET("/:id/likes", userController.GetLikesByID)

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
	blogRouter.PUT("/:id/comments/:comment_id")
	blogRouter.DELETE("/:id/comments/:comment_id")

	blogRouter.GET("/:id/shares")
	blogRouter.POST("/:id/shares")
	blogRouter.DELETE("/:id/shares/:share_id")

	blogRouter.GET("/:id/ratings")
	blogRouter.POST("/:id/ratings")
	blogRouter.PUT("/:id/ratings/:ratings_id")
	blogRouter.DELETE("/:id/ratings/:ratings_id")

	tagRouter := gin.Group("tag")
	tagRouter.Use(middleware.AuthMiddleware(env.JwtSecret))
	tagRouter.GET("/")
	tagRouter.GET("/:id")
	tagRouter.POST("/")
	tagRouter.PUT("/:id")
	tagRouter.DELETE("/:id")
}
