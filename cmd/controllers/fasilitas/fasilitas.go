package controllers

import (
  "net/http"
  m_fasilitas "koskuy-ws/cmd/models/fasilitas"
  "github.com/labstack/echo"
)


func GetAllFasilitas(c echo.Context) error {
  data       := m_fasilitas.GetAllFasilitas()
  return c.JSON(http.StatusOK, data)
}
