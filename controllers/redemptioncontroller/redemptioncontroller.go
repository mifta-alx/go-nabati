package redemptioncontroller

import (
	"github.com/labstack/echo/v4"
	"go-nabati/models"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
)

func Index(ctx echo.Context) error {
	var uniqueCode models.UniqueCode
	random := strconv.Itoa(rand.Intn(1000000))
	uniqueCode.Code = random
	models.DB.Create(&uniqueCode)

	var uniqueCodes []models.UniqueCode
	models.DB.Find(&uniqueCodes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"uniquecodes": uniqueCodes,
	})
}

func Check(ctx echo.Context) error {

	var input struct {
		Code string `json:"code"`
	}
	if err := ctx.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var uniqueCode models.UniqueCode
	if err := models.DB.Model(&uniqueCode).Where("code = ?", input.Code).First(&uniqueCode).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Code Not Found",
			})
		default:
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
	}
	if uniqueCode.Status == 0 {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message":     "Valid Code",
			"status_code": uniqueCode.Status,
		})
	} else {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":     "Invalid Code",
			"status_code": uniqueCode.Status,
		})
	}
}

func Submit(ctx echo.Context) error {
	var userData models.Redemption
	if err := ctx.Bind(&userData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if err := models.DB.Create(&userData).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Redeem successfully",
	})
}
