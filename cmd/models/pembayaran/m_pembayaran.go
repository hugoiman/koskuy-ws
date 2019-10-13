package models

import (
  "fmt"
  "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func GetLaporanPembayaran(id_kos, bulan, tahun string) structs.PembayaranList {
  con     :=  db.Connect()
  query   :=  "SELECT b.nama, b.kamar, b.foto, b.id_kos, a.id_pembayaran, a.id_renter, a.id_member, a.tipe_pembayaran, a.durasi, a.tanggal_masuk, a.tanggal_akhir, a.tanggal_penagihan, a.harga_sewa, a.total_pembayaran, a.jatuh_tempo, a.total_dibayar, a.tagihan, a.status_pembayaran FROM pembayaran a  JOIN renter b ON a.id_renter = b.id_renter WHERE b.id_kos = ?"
  rows, err := con.Query(query, id_kos)

  if err != nil {
    fmt.Println(err.Error())
  }

  pembayaran       := structs.Pembayaran{}
  pembayaran_list  := structs.PembayaranList{}

  for rows.Next() {
    err2 := rows.Scan(
      &pembayaran.Nama, &pembayaran.Kamar, &pembayaran.Foto,
      &pembayaran.Id_kos, &pembayaran.Id_pembayaran, &pembayaran.Id_renter, &pembayaran.Id_member,
      &pembayaran.Tipe_pembayaran, &pembayaran.Durasi,
      &pembayaran.Tanggal_masuk_ori, &pembayaran.Tanggal_akhir_ori, &pembayaran.Tanggal_penagihan_ori,
      &pembayaran.Harga_sewa, &pembayaran.Total_pembayaran, &pembayaran.Jatuh_tempo_ori,
      &pembayaran.Total_dibayar, &pembayaran.Tagihan, &pembayaran.Status_pembayaran,
    )
    pembayaran.Tanggal_masuk = pembayaran.Tanggal_masuk_ori.Format("02 Jan 2006")
    pembayaran.Tanggal_akhir = pembayaran.Tanggal_akhir_ori.Format("02 Jan 2006")
    pembayaran.Tanggal_penagihan = pembayaran.Tanggal_penagihan_ori.Format("02 Jan 2006")
    pembayaran.Jatuh_tempo = pembayaran.Jatuh_tempo_ori.Format("02 Jan 2006")

    pembayaran.BiayaTambahanList = make([]structs.Biaya_tambahan,0)
    pembayaran.TanggalPembayaranList = make([]structs.Tanggal_pembayaran,0)

    if err2 != nil {
      fmt.Println(err2.Error())
    }
    pembayaran_list.PembayaranList = append(pembayaran_list.PembayaranList, pembayaran)
  }

  //  Get Pembayaran Lain
  for key, value := range pembayaran_list.PembayaranList{
    query     :=  "SELECT id_biaya, keterangan, nominal FROM biaya_tambahan WHERE id_pembayaran = ?"
    rows2, _  := con.Query(query, value.Id_pembayaran)
    for rows2.Next() {
      var biaya_tambahan structs.Biaya_tambahan

      _ = rows2.Scan(&biaya_tambahan.Id_biaya, &biaya_tambahan.Keterangan, &biaya_tambahan.Nominal)
      pembayaran_list.PembayaranList[key].BiayaTambahanList = append(pembayaran_list.PembayaranList[key].BiayaTambahanList, biaya_tambahan)
    }
  }

  //  Get Tanggal Pembayaran
  for key2, value2 := range pembayaran_list.PembayaranList{
    query     :=  "SELECT id_tanggal_pembayaran, tanggal_pembayaran, nominal FROM tanggal_pembayaran WHERE id_pembayaran = ?"
    rows2, _  := con.Query(query, value2.Id_pembayaran)
    for rows2.Next() {
      var tanggal_pembayaran structs.Tanggal_pembayaran

      _ = rows2.Scan(&tanggal_pembayaran.Id_tanggal_pembayaran, &tanggal_pembayaran.Tanggal_pembayaran_ori, &tanggal_pembayaran.Nominal)
      tanggal_pembayaran.Tanggal_pembayaran = tanggal_pembayaran.Tanggal_pembayaran_ori.Format("02 Jan 2006")

      pembayaran_list.PembayaranList[key2].TanggalPembayaranList = append(pembayaran_list.PembayaranList[key2].TanggalPembayaranList, tanggal_pembayaran)
    }
  }

  defer con.Close()

  return pembayaran_list
}

func GetPembayaran(id_pembayaran string) structs.Pembayaran {
  con     :=  db.Connect()
  query   :=  "SELECT b.nama, b.kamar, b.foto, b.id_kos, a.id_pembayaran, a.id_renter, a.id_member, a.tipe_pembayaran, a.durasi, a.tanggal_masuk, a.tanggal_akhir, a.tanggal_penagihan, a.harga_sewa, a.total_pembayaran, a.jatuh_tempo, a.total_dibayar, a.tagihan, a.status_pembayaran FROM pembayaran a  JOIN renter b ON a.id_renter = b.id_renter  WHERE a.id_pembayaran = ?"
  rows, err := con.Query(query, id_pembayaran)

  if err != nil {
    fmt.Println(err.Error())
  }

  pembayaran       := structs.Pembayaran{}

  for rows.Next() {
    err2 := rows.Scan(
      &pembayaran.Nama, &pembayaran.Kamar, &pembayaran.Foto,
      &pembayaran.Id_kos, &pembayaran.Id_pembayaran, &pembayaran.Id_renter, &pembayaran.Id_member,
      &pembayaran.Tipe_pembayaran, &pembayaran.Durasi,
      &pembayaran.Tanggal_masuk_ori, &pembayaran.Tanggal_akhir_ori, &pembayaran.Tanggal_penagihan_ori,
      &pembayaran.Harga_sewa, &pembayaran.Total_pembayaran, &pembayaran.Jatuh_tempo_ori,
      &pembayaran.Total_dibayar, &pembayaran.Tagihan, &pembayaran.Status_pembayaran,
    )
    pembayaran.Tanggal_masuk = pembayaran.Tanggal_masuk_ori.Format("02 Jan 2006")
    pembayaran.Tanggal_akhir = pembayaran.Tanggal_akhir_ori.Format("02 Jan 2006")
    pembayaran.Tanggal_penagihan = pembayaran.Tanggal_penagihan_ori.Format("02 Jan 2006")
    pembayaran.Jatuh_tempo = pembayaran.Jatuh_tempo_ori.Format("02 Jan 2006")

    // pembayaran.BiayaTambahanList = make([]structs.Biaya_tambahan, 0)

    if err2 != nil {
      fmt.Println(err2.Error())
    }
  }

  //  Get Pembayaran Lain
  query2    :=  "SELECT id_biaya, keterangan, nominal FROM biaya_tambahan WHERE id_pembayaran = ?"
  rows2, _  := con.Query(query2, pembayaran.Id_pembayaran)

  biaya_tambahan       := structs.Biaya_tambahan{}

  for rows2.Next() {
    _ = rows2.Scan(&biaya_tambahan.Id_biaya, &biaya_tambahan.Keterangan, &biaya_tambahan.Nominal)

    pembayaran.BiayaTambahanList = append(pembayaran.BiayaTambahanList, biaya_tambahan)
  }

  //  Get Tanggal Pembayaran
  query3    :=  "SELECT id_tanggal_pembayaran, tanggal_pembayaran, nominal FROM tanggal_pembayaran WHERE id_pembayaran = ?"
  rows3, _  := con.Query(query3, pembayaran.Id_pembayaran)

  tanggal_pembayaran       := structs.Tanggal_pembayaran{}

  for rows3.Next() {
    _ = rows3.Scan(&tanggal_pembayaran.Id_tanggal_pembayaran, &tanggal_pembayaran.Tanggal_pembayaran_ori, &tanggal_pembayaran.Nominal)

    tanggal_pembayaran.Tanggal_pembayaran = tanggal_pembayaran.Tanggal_pembayaran_ori.Format("02 Jan 2006")
    pembayaran.TanggalPembayaranList = append(pembayaran.TanggalPembayaranList, tanggal_pembayaran)
  }

  defer con.Close()

  return pembayaran
}

func GetHistoryPembayaran(id_renter string) string {
  return id_renter
}
