package controllers

import (
  "net/http"
  "koskuy-ws/cmd/models"
  "github.com/labstack/echo"
)


func GetKosFavorit(c echo.Context) error {
  id_member  := c.Param("id_member")
  data       := models.GetKosFavorit(id_member)
  return c.JSON(http.StatusOK, data)
}
