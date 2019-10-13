package models

import (
  "fmt"
  "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func GetMykosList(id_member string) structs.KosList {
  con     :=  db.Connect()
  query   :=  "SELECT b.id_member, b.otoritas, b.status, a.id_kos, a.nama_kos, a.tipe_kos, a.alamat, a.luas_kamar, a.total_kamar, a.kamar_terisi, a.deskripsi, a.keterangan_lain, a.status_kos, a.booking, a.slug, a.create_at, a.update_at FROM kos a JOIN otoritas b ON a.id_kos = b.id_kos WHERE b.id_member = ? AND b.status = 'aktif'"
  rows, err := con.Query(query, id_member)

  if err != nil {
    fmt.Println(err.Error())
  }

  kos       := structs.Kos{}
  kos_list  := structs.KosList{}

  for rows.Next() {
    err2 := rows.Scan(&kos.Id_member, &kos.Otoritas, &kos.Status,
      &kos.Id_kos, &kos.Nama_kos, &kos.Tipe_kos, &kos.Alamat,
      &kos.Luas_kamar, &kos.Total_kamar, &kos.Kamar_terisi,
      &kos.Deskripsi, &kos.Keterangan_lain, &kos.Status_kos,
      &kos.Booking, &kos.Slug,
      &kos.Create_at_ori, &kos.Update_at_ori,
    )
    kos.Create_at = kos.Create_at_ori.Format("02 Jan 2006")
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

func GetMyKos(slug, id_member string) (structs.Kos,error) {
  con     :=  db.Connect()
  query   :=  "SELECT b.id_member, b.otoritas, b.status, a.id_kos, a.nama_kos, a.tipe_kos, a.alamat, a.luas_kamar, a.total_kamar, a.kamar_terisi, a.deskripsi, a.keterangan_lain, a.status_kos, a.booking, a.slug, a.create_at, a.update_at FROM kos a JOIN otoritas b ON a.id_kos = b.id_kos WHERE a.slug = ? AND b.id_member = ?"

  kos     :=  structs.Kos{}
  err     :=  con.QueryRow(query, slug, id_member).Scan(
      &kos.Id_member, &kos.Otoritas, &kos.Status,
      &kos.Id_kos, &kos.Nama_kos, &kos.Tipe_kos, &kos.Alamat,
      &kos.Luas_kamar, &kos.Total_kamar, &kos.Kamar_terisi,
      &kos.Deskripsi, &kos.Keterangan_lain, &kos.Status_kos,
      &kos.Booking, &kos.Slug,
      &kos.Create_at_ori, &kos.Update_at_ori,
  )
  kos.Create_at = kos.Create_at_ori.Format("02 Jan 2006")
  kos.Update_at = kos.Update_at_ori.Format("02 Jan 2006")
  kos.FotoKosList = make([]structs.Foto_kos,0)

  if err != nil {
    return kos, nil
  }

  //  Get Foto
  query2    := "SELECT id_foto_kos, nama_foto_kos FROM foto_kos WHERE id_kos = ?"
  rows2, _  := con.Query(query2, kos.Id_kos)
  for rows2.Next() {
    var foto_kos structs.Foto_kos

    _ = rows2.Scan(&foto_kos.Id_foto_kos, &foto_kos.Nama_foto_kos)
    kos.FotoKosList = append(kos.FotoKosList, foto_kos)
  }

  //  Get Harga Sewa
  query3    := "SELECT bulanan, harian, mingguan, tahunan FROM harga_sewa WHERE id_kos = ?"
  rows3, _  := con.Query(query3, kos.Id_kos)
  for rows3.Next() {
    var harga_sewa structs.Harga_sewa

    _ = rows3.Scan(&harga_sewa.Bulanan, &harga_sewa.Harian, &harga_sewa.Mingguan, &harga_sewa.Tahunan)
    kos.HargaSewaList = append(kos.HargaSewaList, harga_sewa)
  }

  //  Get Rating
  query4    := "SELECT AVG(kebersihan), AVG(kenyamanan), AVG(keamanan), AVG(fasilitas_kamar), AVG(fasilitas_bersama), AVG(harga), AVG(rating) FROM review WHERE id_kos = ?"
  rows4, _  := con.Query(query4, kos.Id_kos)

  for rows4.Next() {
    var review structs.Review
    _ = rows4.Scan(&review.Kebersihan, &review.Kenyamanan, &review.Keamanan, &review.Fasilitas_kamar, &review.Fasilitas_bersama, &review.Harga, &review.Rating)
    kos.Review = append(kos.Review, review)
  }

  defer con.Close()

  return kos, nil
}
