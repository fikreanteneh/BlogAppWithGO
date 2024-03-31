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
	profileRouter := gin.Group("profile")
	userRouter := gin.Group("user")
	blogRouter := gin.Group("blog")
	tagRouter := gin.Group("tag")


	publicRouter.POST("/register", authController.Register)
	publicRouter.POST("/login", authController.Login)
	publicRouter.POST("/adminRegister", authController.AdminRegister)

	profileRouter.Use(middleware.AuthMiddleware(env.JwtSecret))
	profileRouter.GET("/", profileController.GetProfile)
	profileRouter.PUT("/", profileController.UpdateProfile)
	profileRouter.DELETE("/", profileController.DeleteProfile)
	profileRouter.PUT("/updatePassword", profileController.UpdatePassword) 
	profileRouter.PUT("/updateEmail", profileController.UpdateEmail) 
	profileRouter.PUT("/updateUsername", profileController.UpdateUsername) 
	//TODO: Implement UpdateProfilePicture

	userRouter.GET("/", userController.GetUsers)
	userRouter.GET("/:user_id", userController.GetUserByID)
	userRouter.GET("/:user_id/followers", userController.GetFollowersByID)
	userRouter.POST("/:user_id/follow", middleware.AuthMiddleware(env.JwtSecret), userController.FollowUserByID)
	userRouter.DELETE("/:user_id/follow",middleware.AuthMiddleware(env.JwtSecret), userController.UnfollowUserByID)
	userRouter.GET("/:user_id/followings", userController.GetFollowingsByID)
	userRouter.GET("/:user_id/blogs", userController.GetBlogsByID)
	userRouter.GET("/:user_id/shares", userController.GetSharesByID)
	userRouter.GET("/:user_id/likes", userController.GetLikesByID)

	notificationRouter := gin.Group("notification")
	notificationRouter.GET("/", notificationController.GetNotification)
	notificationRouter.DELETE("/:notification_id", notificationController.DeleteNotificationByID)

	blogRouter.GET("/", blogController.GetAllBlogs)
	blogRouter.GET("/:blog_id", blogController.GetByBlogID)
	blogRouter.POST("/",middleware.AuthMiddleware(env.JwtSecret), blogController.CreateBlog)
	blogRouter.PUT("/:blog_id",middleware.AuthMiddleware(env.JwtSecret), blogController.UpdateBlog)
	blogRouter.DELETE("/:blog_id",middleware.AuthMiddleware(env.JwtSecret), blogController.DeleteBlog)

	blogRouter.GET("/:blog_id/likes", likeController.GetLikesByBlogID)
	blogRouter.POST("/:blog_id/likes", middleware.AuthMiddleware(env.JwtSecret), likeController.LikeBlog)
	blogRouter.DELETE("/:blog_id/likes",middleware.AuthMiddleware(env.JwtSecret), likeController.UnlikeBlog)

	blogRouter.GET("/:blog_id/comments", commentController.GetCommentByBlogID)
	blogRouter.POST("/:blog_id/comments", middleware.AuthMiddleware(env.JwtSecret), commentController.CreateComment)
	blogRouter.PUT("/:blog_id/comments/:comment_id",middleware.AuthMiddleware(env.JwtSecret), commentController.UpdateComment)
	blogRouter.DELETE("/:blog_id/comments/:comment_id",middleware.AuthMiddleware(env.JwtSecret), commentController.DeleteComment)

	blogRouter.GET("/:blog_id/shares", shareController.GetSharesByBlogID)
	blogRouter.POST("/:blog_id/shares",middleware.AuthMiddleware(env.JwtSecret),  shareController.ShareBlog)
	blogRouter.DELETE("/:blog_id/shares/:share_id", middleware.AuthMiddleware(env.JwtSecret), shareController.DeleteShare)

	blogRouter.GET("/:blog_id/ratings", ratingController.GetRatingsByBlogID)
	blogRouter.POST("/:blog_id/ratings", middleware.AuthMiddleware(env.JwtSecret), ratingController.RateBlog)
	blogRouter.PUT("/:blog_id/ratings/:rating_id",middleware.AuthMiddleware(env.JwtSecret), ratingController.UpdateRating)
	//TODO: Implement DeleteBlogRating

	tagRouter.GET("/", tagController.GetTags)
	tagRouter.GET("/:tag_id", tagController.GetTagByID)
	tagRouter.POST("/",middleware.AuthMiddleware(env.JwtSecret), tagController.CreateTag)
	tagRouter.PUT("/:tag_id", middleware.AuthMiddleware(env.JwtSecret),tagController.UpdateTagByID)
	tagRouter.DELETE("/:tag_id", middleware.AuthMiddleware(env.JwtSecret), tagController.DeleteTagByID)
}
