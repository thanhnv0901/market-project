package controllers

import "github.com/labstack/echo/v4"

// Responce ..
func Responce(c echo.Context, statusCode int, message string, isSuccess bool, data interface{}) error {
	return c.JSON(statusCode, map[string]interface{}{
		"success": isSuccess,
		"message": message,
		"data":    data,
	})
}
