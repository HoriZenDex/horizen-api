package controllers

import (
	"net/http"
	"strconv"

	"horizen-api/config"
	"horizen-api/models"

	"github.com/labstack/echo/v4"
)

// CreateComment creates a comment for a given post.
func CreateComment(c echo.Context) error {
	postID, _ := strconv.Atoi(c.Param("id"))
	user := c.Get("user").(*models.User)

	comment := new(models.Comment)
	if err := c.Bind(comment); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	comment.PostID = uint(postID)
	comment.UserID = user.ID

	config.DB.Create(&comment)
	return c.JSON(http.StatusCreated, comment)
}

// DeleteComment deletes a comment by its ID.
func DeleteComment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.Delete(&models.Comment{}, id)
	return c.JSON(http.StatusOK, echo.Map{"message": "Comment deleted"})
}
