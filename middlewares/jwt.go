package middlewares

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		tokenString := c.Request().Header.Get("Authorization")

		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Unauthorized: Token is missing",
			})
		}

		bearerToken := strings.Split(tokenString, "Bearer ")
		if len(bearerToken) != 2 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Unauthorized: Malformed token",
			})
		}

		tokenString = strings.TrimSpace(bearerToken[1])

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("API_SECRET")), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Unauthorized: " + err.Error(),
			})
		}

		// Periksa apakah token valid
		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Unauthorized: Invalid token",
			})
		}

		// Set claims dari token ke konteks untuk digunakan di handler
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Server Error",
			})
		}
		c.Set("user", claims)
		return next(c)
	}
}
