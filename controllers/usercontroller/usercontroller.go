package usercontroller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-nabati/models"
	"gorm.io/gorm"
	"net/http"
)

func Index(ctx echo.Context) error {
	var users []models.User
	models.DB.Find(&users)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

func Show(ctx echo.Context) error {
	var user models.User
	id := ctx.Param("id")
	if err := models.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
				"message": "User not found",
			})
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func Create(ctx echo.Context) error {
	var user models.User

	if err := ctx.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	models.DB.Create(&user)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User created successfully",
	})
}

func Update(ctx echo.Context) error {
	var user models.User
	id := ctx.Param("id")

	if err := ctx.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if models.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "User failed to update",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
	})
}

func Delete(ctx echo.Context) error {
	var user models.User

	var input struct {
		Id json.Number
	}

	if err := ctx.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&user, id).RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "User failed to delete",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}
