package models

import (
  "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func CreateTanggalPembayaran(data structs.AddTanggalPembayaran) bool {
  con     :=  db.Connect()
	query 	:=  "INSERT INTO tanggal_pembayaran (id_pembayaran, tanggal_pembayaran, nominal) VALUES (?,?,?)"
  _, err  :=  con.Exec(query, data.Id_pembayaran, data.Tanggal_pembayaran, data.Nominal)

	defer con.Close()

  if err == nil {
    return true
  } else {
    return false
  }
}
