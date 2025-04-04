package routes

import (
	"horizen-api/controllers"
	"horizen-api/middleware"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes sets up both public and protected routes.
func RegisterRoutes(e *echo.Echo) {
	// Public routes (no JWT required)
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	// Protected API group (requires JWT)
	api := e.Group("/api")
	api.Use(middleware.JWTMiddleware)

	// Post routes
	api.POST("/posts", controllers.CreatePost)
	api.GET("/posts", controllers.GetAllPosts)
	api.GET("/posts/:id", controllers.GetPostByID)
	api.DELETE("/posts/:id", controllers.DeletePost)

	// Comment routes
	api.POST("/posts/:id/comments", controllers.CreateComment)
	api.DELETE("/comments/:id", controllers.DeleteComment)

	// Like routes
	api.POST("/posts/:id/like", controllers.LikePost)
	api.DELETE("/posts/:id/unlike", controllers.UnlikePost)
}
