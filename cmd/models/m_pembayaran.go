package models

import (
  "fmt"
  "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func GetLaporanPembayaran(id_kos, bulan, tahun string) structs.LaporanPembayaranList {
  con     :=  db.Connect()
  query   :=  "SELECT b.nama, b.kamar, b.foto, a.id_pembayaran, a.id_renter, a.tanggal_akhir, a.total_pembayaran, a.tagihan, a.status_pembayaran, c.tanggal_pembayaran, c.nominal FROM pembayaran a JOIN renter b ON a.id_renter = b.id_renter JOIN tanggal_pembayaran c ON a.id_pembayaran = c.id_pembayaran WHERE b.id_kos = ? AND MONTH(c.tanggal_pembayaran) = ? AND YEAR(c.tanggal_pembayaran) = ?"
  rows, err := con.Query(query, id_kos, bulan, tahun)

  if err != nil {
    fmt.Println(err.Error())
  }

  pembayaran       := structs.LaporanPembayaran{}
  pembayaran_list  := structs.LaporanPembayaranList{}

  for rows.Next() {
    err2 := rows.Scan(
      &pembayaran.Nama, &pembayaran.Kamar, &pembayaran.Foto,
      &pembayaran.Id_pembayaran, &pembayaran.Id_renter,
      &pembayaran.Tanggal_akhir_ori,
      &pembayaran.Total_pembayaran, &pembayaran.Tagihan, &pembayaran.Status_pembayaran,
      &pembayaran.Tanggal_pembayaran_ori, &pembayaran.Nominal,
    )

    pembayaran.Tanggal_akhir = pembayaran.Tanggal_akhir_ori.Format("02 Jan 2006")
    pembayaran.Tanggal_pembayaran = pembayaran.Tanggal_pembayaran_ori.Format("02 Jan 2006")

    pembayaran_list.LaporanPembayaranList = append(pembayaran_list.LaporanPembayaranList, pembayaran)

    if err2 != nil {
      fmt.Println(err2.Error())
    }
  }

  defer con.Close()

  return pembayaran_list
}

func GetStatusPembayaran(id_kos string) (int, int, int, int) {
  var total_renter, lunas, angsur, belum_bayar int
  con     :=  db.Connect()

  query1  :=  "SELECT COUNT(a.id_renter) FROM renter a JOIN pembayaran b ON a.id_renter = b.id_renter WHERE a.id_kos = ? AND a.status_renter = 'aktif' AND b.id_pembayaran IN (SELECT MAX(b.id_pembayaran) FROM pembayaran b GROUP BY b.id_renter)"
  query2  :=  "SELECT COUNT(a.id_renter) FROM renter a JOIN pembayaran b ON a.id_renter = b.id_renter WHERE a.id_kos = ? AND a.status_renter = 'aktif' AND b.status_pembayaran = 'lunas' AND b.id_pembayaran IN (SELECT MAX(b.id_pembayaran) FROM pembayaran b GROUP BY b.id_renter)"
  query3  :=  "SELECT COUNT(a.id_renter) FROM renter a JOIN pembayaran b ON a.id_renter = b.id_renter WHERE a.id_kos = ? AND a.status_renter = 'aktif' AND b.status_pembayaran = 'angsur' AND b.id_pembayaran IN (SELECT MAX(b.id_pembayaran) FROM pembayaran b GROUP BY b.id_renter)"
  query4  :=  "SELECT COUNT(a.id_renter) FROM renter a JOIN pembayaran b ON a.id_renter = b.id_renter WHERE a.id_kos = ? AND a.status_renter = 'aktif' AND b.status_pembayaran = 'belum bayar' AND b.id_pembayaran IN (SELECT MAX(b.id_pembayaran) FROM pembayaran b GROUP BY b.id_renter)"

  err1 :=  con.QueryRow(query1, id_kos).Scan(&total_renter)
  err2 :=  con.QueryRow(query2, id_kos).Scan(&lunas)
  err3 :=  con.QueryRow(query3, id_kos).Scan(&angsur)
  err4 :=  con.QueryRow(query4, id_kos).Scan(&belum_bayar)

  if err1 != nil { fmt.Println(err1.Error()) }
  if err2 != nil { fmt.Println(err2.Error()) }
  if err3 != nil { fmt.Println(err3.Error()) }
  if err4 != nil { fmt.Println(err4.Error()) }

  defer con.Close()

  return total_renter, lunas, angsur, belum_bayar

}
// func GetLaporanPembayaran(id_kos, bulan, tahun string) structs.PembayaranList {
//   con     :=  db.Connect()
//   query   :=  "SELECT b.nama, b.kamar, b.foto, b.id_kos, a.id_pembayaran, a.id_renter, a.id_member, a.tipe_pembayaran, a.durasi, a.tanggal_masuk, a.tanggal_akhir, a.tanggal_penagihan, a.harga_sewa, a.total_pembayaran, a.jatuh_tempo, a.total_dibayar, a.tagihan, a.status_pembayaran FROM pembayaran a  JOIN renter b ON a.id_renter = b.id_renter WHERE b.id_kos = ?"
//   rows, err := con.Query(query, id_kos)
//
//   if err != nil {
//     fmt.Println(err.Error())
//   }
//
//   pembayaran       := structs.Pembayaran{}
//   pembayaran_list  := structs.PembayaranList{}
//
//   for rows.Next() {
//     err2 := rows.Scan(
//       &pembayaran.Nama, &pembayaran.Kamar, &pembayaran.Foto,
//       &pembayaran.Id_kos, &pembayaran.Id_pembayaran, &pembayaran.Id_renter, &pembayaran.Id_member,
//       &pembayaran.Tipe_pembayaran, &pembayaran.Durasi,
//       &pembayaran.Tanggal_masuk_ori, &pembayaran.Tanggal_akhir_ori, &pembayaran.Tanggal_penagihan,
//       &pembayaran.Harga_sewa, &pembayaran.Total_pembayaran, &pembayaran.Jatuh_tempo_ori,
//       &pembayaran.Total_dibayar, &pembayaran.Tagihan, &pembayaran.Status_pembayaran,
//     )
//     pembayaran.Tanggal_masuk = pembayaran.Tanggal_masuk_ori.Format("02 Jan 2006")
//     pembayaran.Tanggal_akhir = pembayaran.Tanggal_akhir_ori.Format("02 Jan 2006")
//     pembayaran.Jatuh_tempo = pembayaran.Jatuh_tempo_ori.Format("02 Jan 2006")
//
//     pembayaran.BiayaTambahanList = make([]structs.Biaya_tambahan,0)
//     pembayaran.TanggalPembayaranList = make([]structs.Tanggal_pembayaran,0)
//
//     if err2 != nil {
//       fmt.Println(err2.Error())
//     }
//     pembayaran_list.PembayaranList = append(pembayaran_list.PembayaranList, pembayaran)
//   }
//
//   //  Get Pembayaran Lain
//   for key, value := range pembayaran_list.PembayaranList{
//     query     :=  "SELECT id_biaya, keterangan, nominal FROM biaya_tambahan WHERE id_pembayaran = ?"
//     rows2, _  := con.Query(query, value.Id_pembayaran)
//     for rows2.Next() {
//       var biaya_tambahan structs.Biaya_tambahan
//
//       _ = rows2.Scan(&biaya_tambahan.Id_biaya, &biaya_tambahan.Keterangan, &biaya_tambahan.Nominal)
//       pembayaran_list.PembayaranList[key].BiayaTambahanList = append(pembayaran_list.PembayaranList[key].BiayaTambahanList, biaya_tambahan)
//     }
//   }
//
//   //  Get Tanggal Pembayaran
//   for key2, value2 := range pembayaran_list.PembayaranList{
//     query     :=  "SELECT id_tanggal_pembayaran, tanggal_pembayaran, nominal FROM tanggal_pembayaran WHERE id_pembayaran = ?"
//     rows2, _  := con.Query(query, value2.Id_pembayaran)
//     for rows2.Next() {
//       var tanggal_pembayaran structs.Tanggal_pembayaran
//
//       _ = rows2.Scan(&tanggal_pembayaran.Id_tanggal_pembayaran, &tanggal_pembayaran.Tanggal_pembayaran_ori, &tanggal_pembayaran.Nominal)
//       tanggal_pembayaran.Tanggal_pembayaran = tanggal_pembayaran.Tanggal_pembayaran_ori.Format("02 Jan 2006")
//
//       pembayaran_list.PembayaranList[key2].TanggalPembayaranList = append(pembayaran_list.PembayaranList[key2].TanggalPembayaranList, tanggal_pembayaran)
//     }
//   }
//
//   defer con.Close()
//
//   return pembayaran_list
// }

func GetPembayaran(id_pembayaran string) structs.Pembayaran {
  con     :=  db.Connect()
  query   :=  "SELECT b.nama, b.foto, c.nama_kos, a.id_kos, a.id_pembayaran, a.id_renter, a.id_member, a.kamar, a.tipe_pembayaran, a.durasi, a.tanggal_masuk, a.tanggal_akhir, a.tanggal_penagihan, a.denda, a.jatuh_tempo, a.harga_sewa, a.total_pembayaran, a.total_dibayar, a.tagihan, a.status_pembayaran FROM pembayaran a JOIN renter b ON a.id_renter = b.id_renter JOIN kos c ON b.id_kos = c.id_kos WHERE a.id_pembayaran = ?"
  rows, err := con.Query(query, id_pembayaran)

  if err != nil {
    fmt.Println(err.Error())
  }

  pembayaran       := structs.Pembayaran{}

  for rows.Next() {
    err2 := rows.Scan(
      &pembayaran.Nama, &pembayaran.Foto, &pembayaran.Nama_kos, &pembayaran.Id_kos,
      &pembayaran.Id_pembayaran, &pembayaran.Id_renter, &pembayaran.Id_member,
      &pembayaran.Kamar, &pembayaran.Tipe_pembayaran, &pembayaran.Durasi,
      &pembayaran.Tanggal_masuk_ori, &pembayaran.Tanggal_akhir_ori, &pembayaran.Tanggal_penagihan,
      &pembayaran.Denda, &pembayaran.Jatuh_tempo_ori,
      &pembayaran.Harga_sewa, &pembayaran.Total_pembayaran,
      &pembayaran.Total_dibayar, &pembayaran.Tagihan, &pembayaran.Status_pembayaran,
    )
    pembayaran.Tanggal_masuk = pembayaran.Tanggal_masuk_ori.Format("02 Jan 2006")
    pembayaran.Tanggal_akhir = pembayaran.Tanggal_akhir_ori.Format("02 Jan 2006")
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

func GetHistoryPembayaranRenter(id_renter, id_kos string) structs.HistoryPembayaranList {
  con     :=  db.Connect()
  query   :=  "SELECT b.nama_kos, a.id_pembayaran, a.tanggal_masuk, a.tanggal_akhir, a.total_pembayaran, a.total_dibayar, a.tagihan, a.status_pembayaran FROM pembayaran a JOIN kos b ON a.id_kos = b.id_kos WHERE a.id_renter = ? AND a.id_kos = ? ORDER BY a.id_pembayaran DESC"
  rows, err := con.Query(query, id_renter, id_kos)

  if err != nil {
    fmt.Println(err.Error())
  }

  history       := structs.HistoryPembayaran{}
  history_list  := structs.HistoryPembayaranList{}

  for rows.Next() {
    err2 := rows.Scan(
      &history.Nama_kos, &history.Id_pembayaran, &history.Tanggal_masuk_ori, &history.Tanggal_akhir_ori,
      &history.Total_pembayaran, &history.Total_dibayar, &history.Tagihan, &history.Status_pembayaran,
    )

    if err2 != nil {
      fmt.Println(err2.Error())
    }

    history.Tanggal_masuk = history.Tanggal_masuk_ori.Format("02 Jan 2006")
    history.Tanggal_akhir = history.Tanggal_akhir_ori.Format("02 Jan 2006")
    history_list.HistoryPembayaranList = append(history_list.HistoryPembayaranList, history)
  }

  defer con.Close()

  return history_list
}

func GetHistoryPembayaranMember(id_member string) structs.HistoryPembayaranList {
  con     :=  db.Connect()
  query   :=  "SELECT b.nama_kos, a.id_pembayaran, a.tanggal_masuk, a.tanggal_akhir, a.total_pembayaran, a.total_dibayar, a.tagihan, a.status_pembayaran FROM pembayaran a JOIN kos b ON a.id_kos = b.id_kos WHERE a.id_renter = ? ORDER BY a.id_pembayaran DESC"
  rows, err := con.Query(query, id_member)

  if err != nil {
    fmt.Println(err.Error())
  }

  history       := structs.HistoryPembayaran{}
  history_list  := structs.HistoryPembayaranList{}

  for rows.Next() {
    err2 := rows.Scan(
      &history.Nama_kos, &history.Id_pembayaran, &history.Tanggal_masuk_ori, &history.Tanggal_akhir_ori,
      &history.Total_pembayaran, &history.Total_dibayar, &history.Tagihan, &history.Status_pembayaran,
    )

    if err2 != nil {
      fmt.Println(err2.Error())
    }

    history.Tanggal_masuk = history.Tanggal_masuk_ori.Format("02 Jan 2006")
    history.Tanggal_akhir = history.Tanggal_akhir_ori.Format("02 Jan 2006")
    history_list.HistoryPembayaranList = append(history_list.HistoryPembayaranList, history)
  }

  defer con.Close()

  return history_list
}

func GetLaporanBulanan(id_kos, tahun string) structs.LaporanBulananList {
  con     :=  db.Connect()
  query   :=  "SELECT a.tanggal_pembayaran, SUM(a.nominal) FROM tanggal_pembayaran a JOIN pembayaran b ON a.id_pembayaran = b.id_pembayaran WHERE b.id_kos = ? AND YEAR(a.tanggal_pembayaran) = ? GROUP BY MONTH(a.tanggal_pembayaran)"
  rows, err := con.Query(query, id_kos, tahun)

  if err != nil {
    fmt.Println(err.Error())
  }

  bulanan       := structs.LaporanBulanan{}
  bulanan_list  := structs.LaporanBulananList{}

  for rows.Next() {
    err2 := rows.Scan(
      &bulanan.Periode_ori, &bulanan.Pemasukan,
    )

    if err2 != nil {
      fmt.Println(err2.Error())
    }

    bulanan.Periode = bulanan.Periode_ori.Format("January 2006")
    bulanan_list.LaporanBulananList = append(bulanan_list.LaporanBulananList, bulanan)
  }

  defer con.Close()

  return bulanan_list
}

func CreatePembayaran(data structs.AddPembayaran) (bool, int) {
  con     :=  db.Connect()
  query   :=  "INSERT INTO pembayaran (id_kos, id_renter, id_member, kamar, tipe_pembayaran, durasi, tanggal_masuk, tanggal_akhir, tanggal_penagihan, denda, jatuh_tempo, harga_sewa, total_pembayaran, total_dibayar, tagihan, status_pembayaran) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
  exec, err := con.Exec(query, data.Id_kos, data.Id_renter, data.Id_member, data.Kamar, data.Tipe_pembayaran, data.Durasi, data.Tanggal_masuk, data.Tanggal_akhir, data.Tanggal_penagihan, data.Denda, data.Jatuh_tempo, data.Harga_sewa, data.Total_pembayaran, data.Total_dibayar, data.Tagihan, data.Status_pembayaran)

  defer con.Close()

  if err == nil {
    id_int64, _ := exec.LastInsertId()
    id_pembayaran := int(id_int64)
    return true, id_pembayaran
  } else {
    return false, 0
  }
}

func DeletePembayaran(id_pembayaran int) bool  {
  con     :=  db.Connect()
  query   :=  "Delete FROM pembayaran WHERE id_pembayaran = ?"
  _, err  :=  con.Exec(query, id_pembayaran)

  defer con.Close()

  if err == nil {
    return true
  } else {
    return false
  }
}
