package controllers

import (
  "net/http"
  "koskuy-ws/cmd/models"
  "github.com/labstack/echo"
)


func GetAllFasilitas(c echo.Context) error {
  data       := models.GetAllFasilitas()
  return c.JSON(http.StatusOK, data)
}
