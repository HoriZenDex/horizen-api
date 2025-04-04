package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"horizen-api/config"
	"horizen-api/models"

	"github.com/labstack/echo/v4"
)

// CreatePost creates a new post.
func CreatePost(c echo.Context) error {
	fmt.Print("Hello\nwORLD")
	// Retrieve the user from the custom JWT middleware.
	user := c.Get("user").(*models.User)

	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	// Set the owner ID based on the authenticated user.
	post.OwnerID = user.ID

	config.DB.Create(&post)
	return c.JSON(http.StatusCreated, post)
}

// GetAllPosts retrieves all posts.
func GetAllPosts(c echo.Context) error {
	var posts []models.Post
	config.DB.Find(&posts)
	return c.JSON(http.StatusOK, posts)
}

// GetPostByID retrieves a post by its ID.
func GetPostByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if result := config.DB.First(&post, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Post not found"})
	}
	return c.JSON(http.StatusOK, post)
}

// DeletePost deletes a post by its ID.
func DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.Delete(&models.Post{}, id)
	return c.JSON(http.StatusOK, echo.Map{"message": "Post deleted"})
}
