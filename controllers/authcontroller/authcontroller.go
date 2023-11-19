package authcontroller

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go-nabati/config"
	"go-nabati/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

//func Login(w http.ResponseWriter, r *http.Request, ctx echo.Context) {
//	var userReq models.User
//	var user models.User
//	if err := models.DB.Model(&user).Where("username = ?", userReq.Username).First(&user).Error; err != nil {
//		switch err {
//		case gorm.ErrRecordNotFound:
//			echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
//				"message": "Username or Password is wrong",
//			})
//			return
//		default:
//			echo.NewHTTPError(http.StatusInternalServerError, err.Error())
//			return
//		}
//	}
//
//	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password)); err != nil {
//		echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
//			"message": "Username or Password is wrong",
//		})
//		return
//	}
//
//	expTime := time.Now().Add(time.Minute * 1)
//	claims := &config.JWTClaim{
//		Username: user.Username,
//		RegisteredClaims: jwt.RegisteredClaims{
//			Issuer:    "jwt-go-nabati",
//			ExpiresAt: jwt.NewNumericDate(expTime),
//		},
//	}
//
//	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token, err := generateToken.SignedString([]byte(os.Getenv("API_SECRET")))
//	if err != nil {
//		echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
//			"message": err.Error(),
//		})
//		return
//	}
//
//	http.SetCookie(w, &http.Cookie{
//		Name:     "token",
//		Path:     "/",
//		Value:    token,
//		HttpOnly: true,
//	})
//
//	ctx.JSON(http.StatusOK, map[string]interface{}{
//		"message": "Login Success",
//	})
//}

func Login(c echo.Context) error {
	var userReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}
	var user models.User
	if err := models.DB.Model(&user).Where("username = ?", userReq.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Username or Password is wrong",
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Username or Password is wrong",
		})
	}

	expTime := time.Now().Add(time.Minute * 3)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt-go-nabati",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := generateToken.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// Set the token in the response header
	c.Response().Header().Set("Authorization", "Bearer "+token)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login Success",
	})
}

func Logout(c echo.Context) error {
	// Hapus token dari header Authorization
	c.Response().Header().Del("Authorization")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Logout successful",
	})
}
