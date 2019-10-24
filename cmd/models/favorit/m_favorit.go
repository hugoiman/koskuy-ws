package models

import (
  "fmt"
  "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func GetKosFavorit(id_member string) structs.FavoritList {
  con       :=  db.Connect()
  query     :=  "SELECT a.id_kos, a.id_member, b.nama_kos, b.tipe_kos, b.booking, b.slug, b.update_at FROM favorit a JOIN kos b WHERE a.id_member = ?"
  rows, err := con.Query(query, id_member)

  if err != nil {
    fmt.Println(err.Error())
  }

  favorit       := structs.Favorit{}
  favorit_list  := structs.FavoritList{}

  for rows.Next(){
    err2  :=  rows.Scan(
      &favorit.Id_kos, &favorit.Id_member, &favorit.Nama_kos, &favorit.Tipe_kos, &favorit.Booking, &favorit.Slug, &favorit.Update_at_ori,
    )

    favorit.Update_at = favorit.Update_at_ori.Format("02 Jan 2006")

    if err2 != nil {
      fmt.Println(err2.Error())
    }
    favorit_list.FavoritList = append(favorit_list.FavoritList, favorit)
  }

  //  Get Foto
  for key, value := range favorit_list.FavoritList{
    query     :=  "SELECT id_foto_kos, nama_foto_kos FROM foto_kos WHERE id_kos = ?"
    rows2, _  := con.Query(query, value.Id_kos)
    for rows2.Next() {
      var fotokos structs.Foto_kos

      _ = rows2.Scan(&fotokos.Id_foto_kos, &fotokos.Nama_foto_kos)
      favorit_list.FavoritList[key].FotoKosList = append(favorit_list.FavoritList[key].FotoKosList, fotokos)
    }
  }

  //  Get Harga Sewa
  for key2, value2 := range favorit_list.FavoritList{
    query     :=  "SELECT bulanan, harian, mingguan, tahunan FROM harga_sewa WHERE id_kos = ?"
    rows3, _  := con.Query(query, value2.Id_kos)
    for rows3.Next() {
      var harga_sewa structs.Harga_sewa

      _ = rows3.Scan(&harga_sewa.Bulanan, &harga_sewa.Harian, &harga_sewa.Mingguan, &harga_sewa.Tahunan)
      favorit_list.FavoritList[key2].HargaSewaList = append(favorit_list.FavoritList[key2].HargaSewaList, harga_sewa)
    }
  }

  defer con.Close()

  return favorit_list
}
