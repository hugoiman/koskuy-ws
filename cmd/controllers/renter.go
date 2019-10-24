package controllers

import (
  "net/http"
  "koskuy-ws/cmd/models"
  "github.com/labstack/echo"
)

func GetDaftarRenter(c echo.Context) error {
  id_kos    := c.Param("id_kos")
  data      := models.GetDaftarRenter(id_kos)
  return c.JSON(http.StatusOK, data)
}

func GetRenter(c echo.Context) error {
  slug := c.Param("slug")
  data, _   := models.GetRenter(slug)
  return c.JSON(http.StatusOK, data)
}
