package routers

import (
	"myGram/src/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (routes *route) RouterGroup() http.Handler {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", routes.ctrl.UserRegister)
		userRouter.POST("/login", routes.ctrl.UserLogin)

		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/:userId", routes.ctrl.UpdateUser)
		userRouter.DELETE("/:userId", routes.ctrl.DeleteUser)
	}

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.GET("/", routes.ctrl.GetSocialMedias)
		socialMediaRouter.POST("/", routes.ctrl.StoreSocialMedia)
		socialMediaRouter.PUT("/:socialmediaId", middlewares.AuthorizeSocialMedia(), routes.ctrl.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialmediaId", middlewares.AuthorizeSocialMedia(), routes.ctrl.DeleteSocialMedia)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", routes.ctrl.GetPhotos)
		photoRouter.POST("/", routes.ctrl.StorePhoto)
		photoRouter.PUT("/:photoId", middlewares.AuthorizePhoto(), routes.ctrl.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.AuthorizePhoto(), routes.ctrl.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", routes.ctrl.GetComments)
		commentRouter.POST("/", routes.ctrl.StoreComment)
		commentRouter.PUT("/:commentId", middlewares.AuthorizeComment(), routes.ctrl.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.AuthorizeComment(), routes.ctrl.DeleteComment)
	}
	return router
}
