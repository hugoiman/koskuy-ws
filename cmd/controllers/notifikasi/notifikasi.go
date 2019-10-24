package controllers

import (
  "net/http"
  m_notifikasi "koskuy-ws/cmd/models/notifikasi"
  "github.com/labstack/echo"
)


func GetNotifikasiBookingList(c echo.Context) error {
  id_member  := c.Param("id_member")
  data       := m_notifikasi.GetNotifikasiList(id_member)
  return c.JSON(http.StatusOK, data)
}
