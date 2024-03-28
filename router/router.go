package router

import (
	"BlogApp/config"
	"BlogApp/controller"
	"BlogApp/middleware"
	"BlogApp/repository"
	"BlogApp/usecase"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *config.Environment, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {

	ctx := context.TODO()
	userRepository := repository.NewUserRepository(db, "users")
	blogRepository := repository.NewBlogRepository(db, "blogs")
	tagRepository := repository.NewTagRepository(db, "tags")
	notificationRepository := repository.NewNotificationRepository(db, "notifications")
	shareRepository := repository.NewShareRepository(db, "shares")
	blogTagRepository := repository.NewBlogTagRepository(db, "blogTags")
	commentRepository := repository.NewCommentRepository(db, "comments")
	followRepository := repository.NewFollowRepository(db, "follows")
	likeRepository := repository.NewLikeRepository(db, "likes")
	ratingRepository := repository.NewBlogRatingRepository(db, "ratings")

	authUseCase := usecase.NewAuthUseCase(&ctx, env, &userRepository )
	profileUsecase := usecase.NewProfileUseCase(&ctx, env, &userRepository)
	userUseCase := usecase.NewUserUseCase(&ctx, env, &userRepository, &followRepository, &blogRepository, &shareRepository, &likeRepository, &blogTagRepository)
	blogUseCase := usecase.NewBlogUseCase(&ctx, env, &blogRepository, &userRepository, &shareRepository, &likeRepository, &ratingRepository, &tagRepository, &blogTagRepository)
	tagUseCase := usecase.NewTagUseCase(&ctx, env, &tagRepository)
	notificationUseCase := usecase.NewNotificationUseCase(&ctx, env, &notificationRepository)
	shareUseCase := usecase.NewShareUseCase(&ctx, env, &shareRepository, &blogRepository)
	commentUseCase := usecase.NewCommentUseCase(&ctx, env, &commentRepository)
	likeUseCase := usecase.NewLikeUseCase(&ctx, env, &likeRepository, &userRepository)
	ratingUseCase := usecase.NewRatingUseCase(&ctx, env, &ratingRepository)

	authController := controller.NewAuthController(env, &authUseCase)
	profileController :=  controller.NewProfileController(env, &profileUsecase)
	userController := controller.NewUserController(env, &userUseCase)
	blogController := controller.NewBlogController(env, &blogUseCase)
	tagController := controller.NewTagController(env, &tagUseCase)
	notificationController := controller.NewNotificationController(env, &notificationUseCase)
	shareController := controller.NewShareController(env, &shareUseCase)
	commentController := controller.NewCommentController(env, &commentUseCase)
	likeController := controller.NewLikeController(env, &likeUseCase)
	ratingController := controller.NewRatingController(env, &ratingUseCase)

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
	userRouter.GET("/:user_id", userController.GetUserByID)
	userRouter.GET("/:user_id/followers", userController.GetFollowersByID)
	userRouter.POST("/:user_id/follow", userController.FollowUserByID)
	userRouter.DELETE("/:user_id/follow", userController.UnfollowUserByID)
	userRouter.GET("/:user_id/followings", userController.GetFollowingsByID)
	userRouter.GET("/:user_id/blogs", userController.GetBlogsByID)
	userRouter.GET("/:user_id/shares", userController.GetSharesByID)
	userRouter.GET("/:user_id/likes", userController.GetLikesByID)

	notificationRouter := gin.Group("notification")
	userRouter.Use(middleware.AuthMiddleware(env.JwtSecret))
	notificationRouter.GET("/", notificationController.GetNotification)
	notificationRouter.DELETE("/:notification_id", notificationController.DeleteNotificationByID)

	blogRouter := gin.Group("blog")
	blogRouter.Use(middleware.AuthMiddleware(env.JwtSecret))
	blogRouter.GET("/", blogController.GetAllBlogs)
	blogRouter.GET("/:blog_id", blogController.GetByBlogID)
	blogRouter.POST("/", blogController.CreateBlog)
	blogRouter.PUT("/:blog_id", blogController.UpdateBlog)
	blogRouter.DELETE("/:blog_id", blogController.DeleteBlog)

	blogRouter.GET("/:blog_id/likes", likeController.GetLikesByBlogID)
	blogRouter.POST("/:blog_id/likes", likeController.LikeBlog)
	blogRouter.DELETE("/:blog_id/likes", likeController.UnlikeBlog)

	blogRouter.GET("/:blog_id/comments", commentController.GetCommentByBlogID)
	blogRouter.POST("/:blog_id/comments", commentController.CreateComment)
	blogRouter.PUT("/:blog_id/comments/:comment_id", commentController.UpdateComment)
	blogRouter.DELETE("/:blog_id/comments/:comment_id", commentController.DeleteComment)

	blogRouter.GET("/:blog_id/shares", shareController.GetSharesByBlogID)
	blogRouter.POST("/:blog_id/shares", shareController.ShareBlog)
	blogRouter.DELETE("/:blog_id/shares/:share_id", shareController.DeleteShare)

	blogRouter.GET("/:blog_id/ratings", ratingController.GetRatingsByBlogID)
	blogRouter.POST("/:blog_id/ratings", ratingController.RateBlog)
	blogRouter.PUT("/:blog_id/ratings/:rating_id", ratingController.UpdateRating)
	//TODO: Implement DeleteBlogRating
	// blogRouter.DELETE("/:id/ratings/:ratings_id", blogController.DeleteBlogRating)

	tagRouter := gin.Group("tag")
	tagRouter.Use(middleware.AuthMiddleware(env.JwtSecret))
	tagRouter.GET("/", tagController.GetTags)
	tagRouter.GET("/:tag_id", tagController.GetTagByID)
	tagRouter.POST("/", tagController.CreateTag)
	tagRouter.PUT("/:tag_id", tagController.UpdateTagByID)
	tagRouter.DELETE("/:tag_id", tagController.DeleteTagByID)
}
