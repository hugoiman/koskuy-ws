package renter

import (
  "net/http"
  m_renter "koskuy-ws/cmd/models/renter"
  "github.com/labstack/echo"
)

func GetDaftarRenter(c echo.Context) error {
  id_kos    := c.Param("id_kos")
  data      := m_renter.GetDaftarRenter(id_kos)
  return c.JSON(http.StatusOK, data)
}

func GetRenter(c echo.Context) error {
  slug := c.Param("slug")
  data, _   := m_renter.GetRenter(slug)
  return c.JSON(http.StatusOK, data)
}
