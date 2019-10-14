package pembayaran

import (
  "net/http"
  m_pembayaran "koskuy-ws/cmd/models/pembayaran"
  "github.com/labstack/echo"
)

type M map[string]interface{}

func GetLaporanPembayaran(c echo.Context) error {
  id_kos    := c.Param("id_kos")
  bulan     := c.QueryParam("bulan")
  tahun     := c.QueryParam("tahun")
  data      := m_pembayaran.GetLaporanPembayaran(id_kos, bulan, tahun)
  return c.JSON(http.StatusOK, data)
}

func GetStatusPembayaran(c echo.Context) error {
  id_kos    := c.Param("id_kos")
  total_renter, lunas, angsur, belum_bayar     := m_pembayaran.GetStatusPembayaran(id_kos)
  return c.JSON(http.StatusOK, M{"total_renter": total_renter,
                                 "lunas": lunas,
                                 "angsur": angsur,
                                 "belum_bayar": belum_bayar})
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
