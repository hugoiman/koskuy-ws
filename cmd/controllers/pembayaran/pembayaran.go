package pembayaran

import (
  "net/http"
  m_pembayaran "koskuy-ws/cmd/models/pembayaran"
  "github.com/labstack/echo"
)

func GetLaporanPembayaran(c echo.Context) error {
  id_kos    := c.Param("id_kos")
  bulan     := c.QueryParam("bulan")
  tahun     := c.QueryParam("tahun")
  data      := m_pembayaran.GetLaporanPembayaran(id_kos, bulan, tahun)
  return c.JSON(http.StatusOK, data)
}

func GetPembayaran(c echo.Context) error {
  id_pembayaran := c.Param("id_pembayaran")
  data          := m_pembayaran.GetPembayaran(id_pembayaran)
  return c.JSON(http.StatusOK, data)
}

func GetHistoryPembayaran(c echo.Context) error {
  id_renter := c.Param("id_renter")
  data      := m_pembayaran.GetHistoryPembayaran(id_renter)
  return c.JSON(http.StatusOK, data)
}
