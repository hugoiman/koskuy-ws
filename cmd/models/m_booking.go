package models

import (
  "fmt"
  "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func GetBookingList(id_member string) structs.BookingList {
  con     :=  db.Connect()
  query   :=  "SELECT c.nama_kos, a.id_booking, a.id_kos, a.id_member, a.tipe_pembayaran, a.durasi, a.status_booking, a.tanggal_awal, a.tanggal_akhir FROM booking a JOIN otoritas b ON a.id_kos = b.id_kos JOIN kos c ON a.id_kos = c.id_kos WHERE b.id_member = ? AND status = 'aktif' ORDER BY a.status_booking DESC"
  rows, err := con.Query(query, id_member)

  if err != nil {
    fmt.Println(err.Error())
  }

  booking       := structs.Booking{}
  booking_list  := structs.BookingList{}

  for rows.Next() {
    err2 := rows.Scan(&booking.Nama_kos, &booking.Id_booking, &booking.Id_kos, &booking.Id_member,
      &booking.Tipe_pembayaran, &booking.Durasi, &booking.Status_booking,
      &booking.Tanggal_awal_ori, &booking.Tanggal_akhir_ori,
    )
    booking.Tanggal_awal  = booking.Tanggal_awal_ori.Format("02 Jan 2006")
    booking.Tanggal_akhir = booking.Tanggal_akhir_ori.Format("02 Jan 2006")

    booking.DataMember    = make([]structs.Member,0)

    if err2 != nil {
      fmt.Println(err2.Error())
    }
    booking_list.BookingList = append(booking_list.BookingList, booking)
  }

  //  Get Data Member
  for key, value := range booking_list.BookingList{
    query     :=  "SELECT id_member, nama, email, no_hp, jenis_kelamin, alamat, foto, ktp, tanggal_lahir from member WHERE id_member = ?"
    rows2, _  := con.Query(query, value.Id_member)
    for rows2.Next() {
      var member structs.Member

      _ = rows2.Scan(&member.Id_member, &member.Nama, &member.Email, &member.No_hp,
        &member.Jenis_kelamin, &member.Alamat, &member.Foto, &member.Ktp,
        &member.Tanggal_lahir_ori,
      )

      // member.Tanggal_lahir = member.Tanggal_lahir_ori.Format("02 Jan 2006")

      booking_list.BookingList[key].DataMember = append(booking_list.BookingList[key].DataMember, member)
    }
  }

  defer con.Close()

  return booking_list
}

func UpdateStatusBooking(id_booking, status_booking string) bool {
  con     :=  db.Connect()
  query   :=  "UPDATE booking SET status_booking = ? WHERE id_booking = ?"
  _, err  :=  con.Exec(query, status_booking, id_booking)

  defer con.Close()

  if err == nil {
    return true
  } else {
    return false
  }
}
