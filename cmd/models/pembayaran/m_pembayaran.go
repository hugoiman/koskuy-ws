package models

import (
  "fmt"
  "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func GetLaporanPembayaran(id_kos, bulan, tahun string) structs.PembayaranList {
  con     :=  db.Connect()
  query   :=  "SELECT b.nama, b.kamar, b.foto, b.id_kos, a.id_pembayaran, a.id_renter, a.tipe_pembayaran, a.durasi, a.tanggal_awal, a.tanggal_akhir, a.harga_sewa, a.total, a.jatuh_tempo, a.dibayar, a.tagihan, a.status_pembayaran, a.tanggal_dibayar FROM pembayaran a  JOIN renter b ON a.id_renter = b.id_renter  WHERE b.id_kos = ?"
  rows, err := con.Query(query, id_kos)

  if err != nil {
    fmt.Println(err.Error())
  }

  pembayaran       := structs.Pembayaran{}
  pembayaran_list  := structs.PembayaranList{}

  for rows.Next() {
    err2 := rows.Scan(
      &pembayaran.Nama, &pembayaran.Kamar, &pembayaran.Foto, &pembayaran.Id_kos,
      &pembayaran.Id_pembayaran, &pembayaran.Id_renter, &pembayaran.Tipe_pembayaran, &pembayaran.Durasi,
      &pembayaran.Tanggal_awal_ori, &pembayaran.Tanggal_akhir_ori, &pembayaran.Harga_sewa, &pembayaran.Total, &pembayaran.Jatuh_tempo_ori,
      &pembayaran.Dibayar, &pembayaran.Tagihan, &pembayaran.Status_pembayaran, &pembayaran.Tanggal_dibayar_ori,
    )
    pembayaran.Tanggal_awal = pembayaran.Tanggal_awal_ori.Format("02 Jan 2006")
    pembayaran.Tanggal_akhir = pembayaran.Tanggal_akhir_ori.Format("02 Jan 2006")
    pembayaran.Jatuh_tempo = pembayaran.Jatuh_tempo_ori.Format("02 Jan 2006")
    pembayaran.Tanggal_dibayar = pembayaran.Tanggal_dibayar_ori.Format("02 Jan 2006")

    pembayaran.BiayaTambahanList = make([]structs.Biaya_tambahan,0)

    if err2 != nil {
      fmt.Println(err2.Error())
    }
    pembayaran_list.PembayaranList = append(pembayaran_list.PembayaranList, pembayaran)
  }

  //  Get Pembayaran Lain
  for key, value := range pembayaran_list.PembayaranList{
    query     :=  "SELECT id_biaya, deskripsi, biaya FROM biaya_tambahan WHERE id_pembayaran = ?"
    rows2, _  := con.Query(query, value.Id_pembayaran)
    for rows2.Next() {
      var biaya_tambahan structs.Biaya_tambahan

      _ = rows2.Scan(&biaya_tambahan.Id_biaya, &biaya_tambahan.Deskripsi, &biaya_tambahan.Biaya)
      pembayaran_list.PembayaranList[key].BiayaTambahanList = append(pembayaran_list.PembayaranList[key].BiayaTambahanList, biaya_tambahan)
    }
  }

  defer con.Close()

  return pembayaran_list
}

func GetPembayaran(id_pembayaran string) structs.Pembayaran {
  con     :=  db.Connect()
  query   :=  "SELECT b.nama, b.kamar, b.foto, a.id_pembayaran, a.id_renter, a.tipe_pembayaran, a.durasi, a.tanggal_awal, a.tanggal_akhir, a.harga_sewa, a.total, a.jatuh_tempo, a.dibayar, a.tagihan, a.status_pembayaran, a.tanggal_dibayar FROM pembayaran a  JOIN renter b ON a.id_renter = b.id_renter  WHERE a.id_pembayaran = ?"
  rows, err := con.Query(query, id_pembayaran)

  if err != nil {
    fmt.Println(err.Error())
  }

  pembayaran       := structs.Pembayaran{}

  for rows.Next() {
    err2 := rows.Scan(
      &pembayaran.Nama, &pembayaran.Kamar, &pembayaran.Foto,
      &pembayaran.Id_pembayaran, &pembayaran.Id_renter, &pembayaran.Tipe_pembayaran, &pembayaran.Durasi,
      &pembayaran.Tanggal_awal_ori, &pembayaran.Tanggal_akhir_ori, &pembayaran.Harga_sewa, &pembayaran.Total, &pembayaran.Jatuh_tempo_ori,
      &pembayaran.Dibayar, &pembayaran.Tagihan, &pembayaran.Status_pembayaran, &pembayaran.Tanggal_dibayar_ori,
    )
    pembayaran.Tanggal_awal = pembayaran.Tanggal_awal_ori.Format("02 Jan 2006")
    pembayaran.Tanggal_akhir = pembayaran.Tanggal_akhir_ori.Format("02 Jan 2006")
    pembayaran.Jatuh_tempo = pembayaran.Jatuh_tempo_ori.Format("02 Jan 2006")
    pembayaran.Tanggal_dibayar = pembayaran.Tanggal_dibayar_ori.Format("02 Jan 2006")

    // pembayaran.BiayaTambahanList = make([]structs.Biaya_tambahan, 0)

    if err2 != nil {
      fmt.Println(err2.Error())
    }
  }

  //  Get Pembayaran Lain
  query2    :=  "SELECT id_biaya, deskripsi, biaya FROM biaya_tambahan WHERE id_pembayaran = ?"
  rows2, _  := con.Query(query2, pembayaran.Id_pembayaran)

  biaya_tambahan       := structs.Biaya_tambahan{}

  for rows2.Next() {
    _ = rows2.Scan(&biaya_tambahan.Id_biaya, &biaya_tambahan.Deskripsi, &biaya_tambahan.Biaya)

    pembayaran.BiayaTambahanList = append(pembayaran.BiayaTambahanList, biaya_tambahan)
  }
  fmt.Println(biaya_tambahan)


  defer con.Close()

  return pembayaran
}

func GetHistoryPembayaran(id_renter string) string {
  return id_renter
}
