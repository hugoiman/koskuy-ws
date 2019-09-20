package structs

import(
  "time"
)

type PembayaranList struct {
  PembayaranList        []Pembayaran        `json:"pembayaran_list"`
}

type Pembayaran struct {
  Nama                  string              `json:"nama"`
  Kamar                 string              `json:"kamar"`
  Foto                  string              `json:"foto"`
  Id_kos                int                 `json:"id_kos"`

  Id_pembayaran         int                 `json:"id_pembayaran"`
  Id_renter             int                 `json:"id_renter"`
  Tipe_pembayaran       string              `json:"tipe_pembayaran"`
  Durasi                string              `json:"durasi"`
  Tanggal_awal_ori      time.Time           `json:"tanggal_awal_ori"`
  Tanggal_akhir_ori     time.Time           `json:"tanggal_akhir_ori"`
  Harga_sewa            int                 `json:"harga_sewa"`
  Total                 int                 `json:"total"`
  Jatuh_tempo_ori       time.Time           `json:"jatuh_tempo_ori"`
  Dibayar               int                 `json:"dibayar"`
  Tagihan               int                 `json:"tagihan"`
  Status_pembayaran     string              `json:"status_pembayaran"`
  Tanggal_dibayar_ori   time.Time           `json:"tanggal_dibayar_ori"`

  Tanggal_awal          string              `json:"tanggal_awal"`
  Tanggal_akhir         string              `json:"tanggal_akhir"`
  Jatuh_tempo           string              `json:"jatuh_tempo"`
  Tanggal_dibayar       string              `json:"tanggal_dibayar"`

  PembayaranLainList    []Pembayaran_lain   `json:"pembayaran_lain"`
}

type Pembayaran_lain struct {
  Id_pembayaran_lain int                 `json:"id_pembayaran_lain"`
  Deskripsi          string              `json:"deskripsi"`
  Jumlah             int                 `json:"jumlah"`
}
