package models

import (
  "fmt"
  "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func GetNotifikasiList(id_member string) structs.NotifikasiList {
  con     :=  db.Connect()
  query   :=  "SELECT id_notifikasi, id_member, subjek, pesan, status_notif, date FROM notifikasi WHERE id_member = ?"
  rows, err := con.Query(query, id_member)

  if err != nil {
    fmt.Println(err.Error())
  }

  notifikasi       := structs.Notifikasi{}
  notifikasi_list  := structs.NotifikasiList{}

  for rows.Next() {
    err2 := rows.Scan(&notifikasi.Id_notifikasi, &notifikasi.Id_member,
      &notifikasi.Subjek, &notifikasi.Pesan, &notifikasi.Status_notif,
      &notifikasi.Date_ori,
    )
    notifikasi.Date  = notifikasi.Date_ori.Format("02 Jan 2006")

    if err2 != nil {
      fmt.Println(err2.Error())
    }
    notifikasi_list.NotifikasiList = append(notifikasi_list.NotifikasiList, notifikasi)
  }
  defer con.Close()

  return notifikasi_list
}
