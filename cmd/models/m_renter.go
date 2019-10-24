package models

import (
  "fmt"
  "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func GetRenter(slug string) (structs.Renter, error) {
  con     :=  db.Connect()
  query   :=  "SELECT id_renter, id_kos, id_member, nama, email, no_hp, jenis_kelamin, alamat, pekerjaan, foto, ktp, kamar, status_renter, slug, tanggal_lahir from renter WHERE slug = ?"

  renter  :=  structs.Renter{}
  err     :=  con.QueryRow(query, slug).Scan(
    &renter.Id_renter, &renter.Id_kos, &renter.Id_member, &renter.Nama, &renter.Email, &renter.No_hp, &renter.Jenis_kelamin,
    &renter.Alamat, &renter.Pekerjaan, &renter.Foto, &renter.Ktp, &renter.Kamar, &renter.Status_renter, &renter.Slug, &renter.Tanggal_lahir_ori,
  )
  renter.Tanggal_lahir = renter.Tanggal_lahir_ori.Format("02 Jan 2006")

  if err != nil {
    return renter, err
  }
  defer con.Close()

  return renter, nil
}

func GetDaftarRenter(id_kos string) structs.DaftarRenter {
  con     :=  db.Connect()
  query   :=  "SELECT a.id_renter, a.id_member, a.nama, a.no_hp, a.pekerjaan, a.foto, a.kamar, a.status_renter, a.slug, b.status_pembayaran FROM renter a JOIN pembayaran b ON a.id_renter = b.id_renter WHERE a.id_kos = ? AND b.id_pembayaran IN (SELECT MAX(b.id_pembayaran) FROM pembayaran b GROUP BY b.id_renter)"
  rows, err := con.Query(query, id_kos)

  if err != nil {
    fmt.Println(err.Error())
  }

  renter       := structs.Renters{}
  renter_list  := structs.DaftarRenter{}

  for rows.Next() {
    err2 := rows.Scan(
      &renter.Id_renter, &renter.Id_member, &renter.Nama, &renter.No_hp, &renter.Pekerjaan,
      &renter.Foto, &renter.Kamar, &renter.Status_renter, &renter.Slug, &renter.Status_pembayaran,
    )

    if err2 != nil {
      fmt.Println(err2.Error())
    }
    renter_list.DaftarRenter = append(renter_list.DaftarRenter, renter)
  }

  defer con.Close()

  return renter_list
}
