package middleware

import (
	"net/http"
	"os"
	"strings"

	"horizen-api/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware validates JWT tokens and attaches a simplified user to the context.
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing token"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token format"})
		}
		tokenString := parts[1]

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid or expired token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token claims"})
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token user_id"})
		}

		// Attach a simplified user to the context.
		// For more complex apps, you might query the database to get full user details.
		c.Set("user", &models.User{ID: uint(userIDFloat)})

		return next(c)
	}
}
