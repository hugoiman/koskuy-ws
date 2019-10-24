package structs

import(
  "time"
)

type BookingList struct {
  BookingList          []Booking           `json:"booking_list"`
}

type Booking struct {
  // Nama                 string              `json:"nama"`
  // Foto                 string              `json:"foto"`
  Nama_kos             string              `json:"nama_kos"`

  Id_booking           int                 `json:"id_booking"`
  Id_kos               int                 `json:"id_kos"`
  Id_member            int                 `json:"id_member"`
  Tipe_pembayaran      string              `json:"tipe_pembayaran"`
  Durasi               string              `json:"durasi"`
  Tanggal_awal_ori     time.Time           `json:"tanggal_awal_ori"`
  Tanggal_akhir_ori    time.Time           `json:"tanggal_akhir_ori"`
  Status_booking       string              `json:"status_booking"`

  Tanggal_awal         string              `json:"tanggal_awal"`
  Tanggal_akhir        string              `json:"tanggal_akhir"`

  DataMember           []Member            `json:"data_member"`
}
