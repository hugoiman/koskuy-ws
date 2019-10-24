package controllers

import (
  "net/http"
  "koskuy-ws/cmd/models"
  "github.com/labstack/echo"
  "fmt"
  "koskuy-ws/cmd/structs"
  "encoding/json"
)

type M map[string]interface{}

func GetLaporanPembayaran(c echo.Context) error {
  id_kos    := c.Param("id_kos")
  bulan     := c.QueryParam("bulan")
  tahun     := c.QueryParam("tahun")
  data      := models.GetLaporanPembayaran(id_kos, bulan, tahun)
  return c.JSON(http.StatusOK, data)
}

func GetLaporanBulanan(c echo.Context) error {
  id_kos    := c.Param("id_kos")
  tahun     := c.QueryParam("tahun")
  data      := models.GetLaporanBulanan(id_kos, tahun)
  return c.JSON(http.StatusOK, data)
}

func GetStatusPembayaran(c echo.Context) error {
  id_kos    := c.Param("id_kos")
  total_renter, lunas, angsur, belum_bayar     := models.GetStatusPembayaran(id_kos)
  return c.JSON(http.StatusOK, M{"total_renter": total_renter,
                                 "lunas": lunas,
                                 "angsur": angsur,
                                 "belum_bayar": belum_bayar})
}

func GetPembayaran(c echo.Context) error {
  id_pembayaran := c.Param("id_pembayaran")
  data          := models.GetPembayaran(id_pembayaran)
  return c.JSON(http.StatusOK, data)
}

func GetHistoryPembayaran(c echo.Context) error {
  id_member := c.Param("id_member")
  data      := models.GetHistoryPembayaran(id_member)
  return c.JSON(http.StatusOK, data)
}

func AddPembayaran(c echo.Context) error {
  decoder     := json.NewDecoder(c.Request().Body)
  pembayaran  := structs.AddPembayaran{}
  tanggal_pembayaran   := structs.AddTanggalPembayaran{}
  err         := decoder.Decode(&pembayaran)

  if err != nil {
    fmt.Println(err.Error())
  }

  // fmt.Printf("%+v\n", pembayaran)
  status, id_pembayaran  := models.CreatePembayaran(pembayaran)
  if status == true {
    tanggal_pembayaran.Id_pembayaran = id_pembayaran
    tanggal_pembayaran.Tanggal_pembayaran = pembayaran.Tanggal_pembayaran
    tanggal_pembayaran.Nominal = pembayaran.Total_dibayar

    status2   := models.CreateTanggalPembayaran(tanggal_pembayaran)

    if status2 == true {
      return c.JSON(http.StatusOK, M{"status": status2, "id_pembayaran": id_pembayaran})
    } else if status2 == false {
      models.DeletePembayaran(id_pembayaran)
      return c.JSON(http.StatusOK, M{"status": status2, "id_pembayaran": id_pembayaran})
    }

  } else {
    return c.JSON(http.StatusOK, M{"status": status, "id_pembayaran": id_pembayaran})
  }
  return c.JSON(http.StatusOK, M{"status": status, "id_pembayaran": id_pembayaran})
}
