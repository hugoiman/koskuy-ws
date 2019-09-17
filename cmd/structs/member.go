package structs

import(
  "time"
)

type MemberList struct {
  MemberList        []Member            `json:"member_list"`
}

type Member struct {
  Id_member         int                 `json:"id_member"`
  Username          string              `json:"username"`
  Nama              string              `json:"nama"`
  Email             string              `json:"email"`
  No_hp             string              `json:"no_hp"`
  Jenis_kelamin     string              `json:"jenis_kelamin"`
  Alamat            string              `json:"alamat"`
  Foto              string              `json:"foto"`
  Ktp               string              `json:"ktp"`
  Tanggal_lahir_ori time.Time           `json:"tanggal_lahir_ori"`
  Verifikasi_email  bool                `json:"verifikasi_email"`
  Verifikasi_password bool              `json:"verifikasi_password"`
  Verifikasi_no_hp  bool                `json:"verifikasi_no_hp"`

  Tanggal_lahir     string              `json:"tanggal_lahir"`
}
