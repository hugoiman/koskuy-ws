package models

import (
  "fmt"
  "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func GetRenter(id_renter string) (structs.Renter, error) {
  con     :=  db.Connect()
  query   :=  "SELECT id_renter, id_kos, nama, email, no_hp, jenis_kelamin, alamat, foto, ktp, kamar, status_renter, tanggal_lahir from renter WHERE id_renter = ?"

  renter  :=  structs.Renter{}
  err     :=  con.QueryRow(query, id_renter).Scan(
    &renter.Id_renter, &renter.Id_kos, &renter.Nama, &renter.Email, &renter.No_hp, &renter.Jenis_kelamin,
    &renter.Alamat, &renter.Foto, &renter.Ktp, &renter.Kamar, &renter.Status_renter, &renter.Tanggal_lahir_ori,
  )
  renter.Tanggal_lahir = renter.Tanggal_lahir_ori.Format("02 Jan 2006")

  if err != nil {
    return renter, err
  }
  defer con.Close()

  return renter, nil
}

func GetDaftarRenter(id_kos string) structs.RenterList {
  con     :=  db.Connect()
  query   :=  "SELECT id_renter, id_kos, nama, email, no_hp, jenis_kelamin, alamat, foto, ktp, kamar, status_renter, tanggal_lahir from renter WHERE id_kos = ?"
  rows, err := con.Query(query, id_kos)

  if err != nil {
    fmt.Println(err.Error())
  }

  renter       := structs.Renter{}
  renter_list  := structs.RenterList{}

  for rows.Next() {
    err2 := rows.Scan(
      &renter.Id_renter, &renter.Id_kos, &renter.Nama, &renter.Email, &renter.No_hp, &renter.Jenis_kelamin,
      &renter.Alamat, &renter.Foto, &renter.Ktp, &renter.Kamar, &renter.Status_renter, &renter.Tanggal_lahir_ori,
    )
    renter.Tanggal_lahir = renter.Tanggal_lahir_ori.Format("02 Jan 2006")

    if err2 != nil {
      fmt.Println(err2.Error())
    }
    renter_list.RenterList = append(renter_list.RenterList, renter)
  }

  defer con.Close()

  return renter_list
}
