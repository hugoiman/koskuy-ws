package structs

import(
  "time"
)
type RenterList struct {
  RenterList        []Renter            `json:"renter_list"`
}

type Renter struct {
  Id_renter         int                 `json:"id_renter"`
  Id_kos            int                 `json:"id_kos"`
  Id_member         int                 `json:"id_member"`
  Nama              string              `json:"nama"`
  Email             string              `json:"email"`
  No_hp             string              `json:"no_hp"`
  Jenis_kelamin     string              `json:"jenis_kelamin"`
  Alamat            string              `json:"alamat"`
  Pekerjaan         string              `json:"pekerjaan"`
  Foto              string              `json:"foto"`
  Ktp               string              `json:"ktp"`
  Kamar             string              `json:"kamar"`
  Status_renter     string              `json:"status_renter"`
  Slug              string              `json:"slug"`
  Status_pembayaran string              `json:"status_pembayaran"`
  Tanggal_lahir_ori time.Time           `json:"tanggal_lahir_ori"`

  Tanggal_lahir     string              `json:"tanggal_lahir"`
}
