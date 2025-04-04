package controllers

import (
	"net/http"
	"strconv"

	"horizen-api/config"
	"horizen-api/models"

	"github.com/labstack/echo/v4"
)

// LikePost registers a like for a post.
func LikePost(c echo.Context) error {
	postID, _ := strconv.Atoi(c.Param("id"))
	user := c.Get("user").(*models.User)

	like := models.Like{
		PostID: uint(postID),
		UserID: user.ID,
	}

	config.DB.Create(&like)
	return c.JSON(http.StatusCreated, echo.Map{"message": "Post liked"})
}

// UnlikePost removes a like from a post.
func UnlikePost(c echo.Context) error {
	postID, _ := strconv.Atoi(c.Param("id"))
	user := c.Get("user").(*models.User)

	config.DB.Where("post_id = ? AND user_id = ?", postID, user.ID).Delete(&models.Like{})
	return c.JSON(http.StatusOK, echo.Map{"message": "Like removed"})
}
