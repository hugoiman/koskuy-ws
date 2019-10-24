package models

import (
  "fmt"
  "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func GetAllFasilitas() structs.AllFasilitas {
  con       :=  db.Connect()
  query     :=  "SELECT id_fasilitas, jenis_fasilitas, nama_fasilitas FROM fasilitas"
  rows, err := con.Query(query)

  if err != nil {
    fmt.Println(err.Error())
  }

  fasilitas     := structs.Fasilitas{}
  all_fasilitas := structs.AllFasilitas{}

  for rows.Next(){
    err2  :=  rows.Scan(
      &fasilitas.Id_fasilitas, &fasilitas.Jenis_fasilitas, &fasilitas.Nama_fasilitas,
    )

    if err2 != nil {
      fmt.Println(err2.Error())
    }
    all_fasilitas.AllFasilitas = append(all_fasilitas.AllFasilitas, fasilitas)
  }

  defer con.Close()

  return all_fasilitas
}
