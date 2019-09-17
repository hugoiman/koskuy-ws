package renter

import (
  "net/http"
  m_renter "koskuy-ws/cmd/models/renter"
  "github.com/labstack/echo"
)

func GetDaftarRenter(c echo.Context) error {
  id_kos    := c.QueryParam("kos")
  bulan     := c.QueryParam("bulan")
  tahun     := c.QueryParam("tahun")
  data      := m_renter.GetDaftarRenter(id_kos, bulan, tahun)
  return c.JSON(http.StatusOK, data)
}
