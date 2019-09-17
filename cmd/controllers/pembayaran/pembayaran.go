package pembayaran

import (
  "net/http"
  m_pembayaran "koskuy-ws/cmd/models/pembayaran"
  "github.com/labstack/echo"
)

func GetLaporanPembayaran(c echo.Context) error {
  id_kos    := c.QueryParam("kos")
  bulan     := c.QueryParam("bulan")
  tahun     := c.QueryParam("tahun")
  data      := m_pembayaran.GetLaporanPembayaran(id_kos, bulan, tahun)
  return c.JSON(http.StatusOK, data)
}
