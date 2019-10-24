package controllers

import (
  "net/http"
  m_favorit "koskuy-ws/cmd/models/favorit"
  "github.com/labstack/echo"
)


func GetKosFavorit(c echo.Context) error {
  id_member  := c.Param("id_member")
  data       := m_favorit.GetKosFavorit(id_member)
  return c.JSON(http.StatusOK, data)
}
