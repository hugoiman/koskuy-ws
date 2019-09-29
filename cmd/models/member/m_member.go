package member

import (
  db "koskuy-ws/db"
  "koskuy-ws/cmd/structs"
)

func GetMember(id string) (structs.Member, error) {
  con     :=  db.Connect()
  query   :=  "SELECT id_member, username, nama, email, no_hp, jenis_kelamin, alamat, foto, ktp, verifikasi_email, verifikasi_no_hp, tanggal_lahir from member WHERE id_member = ? OR username = ?"

  member  :=  structs.Member{}
  err     :=  con.QueryRow(query, id, id).Scan(
    &member.Id_member, &member.Username, &member.Nama, &member.Email, &member.No_hp,
    &member.Jenis_kelamin, &member.Alamat, &member.Foto, &member.Ktp,
    &member.Verifikasi_email, &member.Verifikasi_no_hp, &member.Tanggal_lahir_ori,
  )
  member.Tanggal_lahir = member.Tanggal_lahir_ori.Format("02 Jan 2006")

  if err != nil {
    return member, err
  }
  defer con.Close()

  return member, nil
}

func CheckOldPassword(id_member, password_lama string) bool {
  var isValid string
  con     :=  db.Connect()
  query   :=  "SELECT username FROM member WHERE id_member = ? AND password = ?"
  err     :=  con.QueryRow(query, id_member, password_lama).Scan(&isValid)

  defer con.Close()

  if err == nil {
    return true
  } else {
    return false
  }
}

func UpdatePassword(id, password_baru string) bool {
  con     :=  db.Connect()
  query   :=  "UPDATE member SET password = ? WHERE id_member = ? OR email = ?"
  _, err  :=  con.Exec(query, password_baru, id, id)

  defer con.Close()

  if err == nil {
    return true
  } else {
    return false
  }
}

func UpdateMember(id_member int, data structs.Member) bool {
  con     :=  db.Connect()
  query   :=  "UPDATE member SET nama = ?, username = ?, tanggal_lahir = ?, jenis_kelamin= ?, alamat = ? WHERE id_member = ?"
  _, err  :=  con.Exec(query, data.Nama, data.Username, data.Tanggal_lahir, data.Jenis_kelamin, data.Alamat, id_member)

  defer con.Close()

  if err == nil {
    return true
  } else {
    return false
  }
}
