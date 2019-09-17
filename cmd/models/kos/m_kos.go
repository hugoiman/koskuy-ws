package models

import (
  "fmt"
  "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func GetMykosList(id_member string) structs.KosList {
  con     :=  db.Connect()
  query   :=  "SELECT id_kos, id_member, nama_kos, tipe_kos, alamat, luas_kamar, total_kamar, kamar_terisi, deskripsi, verifikasi_kos, update_at FROM kos WHERE id_member = ?"
  rows, err := con.Query(query, id_member)

  if err != nil {
    fmt.Println(err.Error())
  }

  kos       := structs.Kos{}
  kos_list  := structs.KosList{}

  for rows.Next() {
    err2 := rows.Scan(&kos.Id_kos, &kos.Id_member, &kos.Nama_kos, &kos.Tipe_kos, &kos.Alamat,
      &kos.Luas_kamar, &kos.Total_kamar, &kos.Kamar_terisi,
      &kos.Deskripsi, &kos.Verifikasi_kos,
      &kos.Update_at_ori,
    )
    kos.Update_at = kos.Update_at_ori.Format("02 Jan 2006")
    kos.FotoKosList = make([]structs.Foto_kos,0)

    if err2 != nil {
      fmt.Println(err2.Error())
    }
    kos_list.KosList = append(kos_list.KosList, kos)
  }

  //  Get Foto
  for key, value := range kos_list.KosList{
    query     :=  "SELECT id_foto_kos, nama_foto_kos FROM foto_kos WHERE id_kos = ?"
    rows2, _  := con.Query(query, value.Id_kos)
    for rows2.Next() {
      var fotokos structs.Foto_kos

      _ = rows2.Scan(&fotokos.Id_foto_kos, &fotokos.Nama_foto_kos)
      kos_list.KosList[key].FotoKosList = append(kos_list.KosList[key].FotoKosList, fotokos)
    }
  }

  //  Get Harga Sewa
  for key2, value2 := range kos_list.KosList{
    query     :=  "SELECT bulanan, harian, mingguan, tahunan FROM harga_sewa WHERE id_kos = ?"
    rows3, _  := con.Query(query, value2.Id_kos)
    for rows3.Next() {
      var harga_sewa structs.Harga_sewa

      _ = rows3.Scan(&harga_sewa.Bulanan, &harga_sewa.Harian, &harga_sewa.Mingguan, &harga_sewa.Tahunan)
      kos_list.KosList[key2].HargaSewaList = append(kos_list.KosList[key2].HargaSewaList, harga_sewa)
    }
  }

  defer con.Close()

  return kos_list
}

func GetMyKos(id_kos, id_member string) (structs.Kos,error) {
  con     :=  db.Connect()
  query   :=  "SELECT id_kos, id_member, nama_kos, tipe_kos, alamat, luas_kamar, total_kamar, kamar_terisi, deskripsi, verifikasi_kos, update_at FROM kos WHERE id_kos = ? AND id_member = ?"

  kos     :=  structs.Kos{}
  err     :=  con.QueryRow(query, id_kos, id_member).Scan(
      &kos.Id_kos, &kos.Id_member, &kos.Nama_kos, &kos.Tipe_kos, &kos.Alamat,
      &kos.Luas_kamar, &kos.Total_kamar, &kos.Kamar_terisi,
      &kos.Deskripsi, &kos.Verifikasi_kos,
      &kos.Update_at_ori,
  )
  kos.Update_at = kos.Update_at_ori.Format("02 Jan 2006")
  kos.FotoKosList = make([]structs.Foto_kos,0)

  if err != nil {
    return kos, nil
  }

  //  Get Foto
  query2    := "SELECT id_foto_kos, nama_foto_kos FROM foto_kos WHERE id_kos = ?"
  rows2, _  := con.Query(query2, id_kos)
  for rows2.Next() {
    var foto_kos structs.Foto_kos

    _ = rows2.Scan(&foto_kos.Id_foto_kos, &foto_kos.Nama_foto_kos)
    kos.FotoKosList = append(kos.FotoKosList, foto_kos)
  }

  //  Get Harga Sewa
  query3    := "SELECT bulanan, harian, mingguan, tahunan FROM harga_sewa WHERE id_kos = ?"
  rows3, _  := con.Query(query3, id_kos)
  for rows3.Next() {
    var harga_sewa structs.Harga_sewa

    _ = rows3.Scan(&harga_sewa.Bulanan, &harga_sewa.Harian, &harga_sewa.Mingguan, &harga_sewa.Tahunan)
    kos.HargaSewaList = append(kos.HargaSewaList, harga_sewa)
  }

  defer con.Close()

  return kos, nil
}
