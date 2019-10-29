package controllers

// import (
//   "net/http"
//   "koskuy-ws/cmd/models"
//   "github.com/labstack/echo"
//   "koskuy-ws/cmd/structs"
//   "encoding/json"
//   "fmt"
// )
//
// func AddTanggalPembayaran(c echo.Context) error {
//   decoder     := json.NewDecoder(c.Request().Body)
//   tanggal_pembayaran  := structs.AddTanggalPembayaran{}
//
//   err         := decoder.Decode(&tanggal_pembayaran)
//
//   if err != nil {
//     return c.JSON(http.StatusOK, M{"status": false})
//   }
//   fmt.Printf("%+v\n", tanggal_pembayaran)
//   getPembayaran :=  models.GetPembayaran(tanggal_pembayaran.Id_pembayaran)
//
//   if tanggal_pembayaran.Nominal > getPembayaran.Tagihan {
//     return c.JSON(http.StatusOK, M{"status": false, "pesan": "Nominal pembayaran melebihi jumlah tagihan."})
//   } else if tanggal_pembayaran.Nominal == getPembayaran.Tagihan {
//     status_pembayaran := "Lunas"
//     // Update Pembayaran
//   } else if tanggal_pembayaran.Nominal < getPembayaran.Tagihan {
//     status_pembayaran := "Angsur"
//   }
//
//   total_dibayar     := getPembayaran.Total_dibayar + tanggal_pembayaran.Nominal
//   tagihan           := getPembayaran.Tagihan - tanggal_pembayaran.Nominal
//   //  Create Tanggal Pembayaran
//   //  Update Pembayaran
//
//   // status   := models.CreateTanggalPembayaran(tanggal_pembayaran)
//   return c.JSON(http.StatusOK, M{"status": true})
// }
