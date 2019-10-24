package structs

import(
  "time"
)

type PembayaranList struct {
  PembayaranList        []Pembayaran        `json:"pembayaran_list"`
}

type BulananList struct {
  BulananList           []Tanggal_pembayaran `json:"bulanan_list"`
}

type Pembayaran struct {
  Nama                  string              `json:"nama"`
  Kamar                 string              `json:"kamar"`
  Foto                  string              `json:"foto"`
  Nama_kos              string              `json:"nama_kos"`
  Id_kos                int                 `json:"id_kos"`

  Id_pembayaran         int                 `json:"id_pembayaran"`
  Id_renter             int                 `json:"id_renter"`
  Id_member             int                 `json:"id_member"`
  Tipe_pembayaran       string              `json:"tipe_pembayaran"`
  Durasi                string              `json:"durasi"`
  Tanggal_masuk_ori     time.Time           `json:"tanggal_masuk_ori"`
  Tanggal_akhir_ori     time.Time           `json:"tanggal_akhir_ori"`
  Tanggal_penagihan     string              `json:"tanggal_penagihan"`
  Harga_sewa            int                 `json:"harga_sewa"`
  Total_pembayaran      int                 `json:"total_pembayaran"`
  Jatuh_tempo_ori       time.Time           `json:"jatuh_tempo_ori"`
  Total_dibayar         int                 `json:"total_dibayar"`
  Tagihan               int                 `json:"tagihan"`
  Status_pembayaran     string              `json:"status_pembayaran"`

  Tanggal_masuk         string              `json:"tanggal_masuk"`
  Tanggal_akhir         string              `json:"tanggal_akhir"`
  Jatuh_tempo           string              `json:"jatuh_tempo"`

  Nominal                int                `json:"nominal"`
  Tanggal_pembayaran_ori time.Time          `json:"tanggal_pembayaran_ori"`
  Tanggal_pembayaran     string             `json:"tanggal_pembayaran"`


  TanggalPembayaranList []Tanggal_pembayaran `json:"tanggal_pembayaran_list"`
  BiayaTambahanList     []Biaya_tambahan     `json:"biaya_tambahan_list"`
}

type Tanggal_pembayaran struct {
  Id_tanggal_pembayaran int                 `json:"id_tanggal_pembayaran"`
  Tanggal_pembayaran_ori time.Time          `json:"tanggal_pembayaran_ori"`
  Nominal               int                 `json:"nominal"`

  Tanggal_pembayaran    string              `json:"tanggal_pembayaran"`
}

type Biaya_tambahan struct {
  Id_biaya              int                 `json:"id_biaya"`
  Keterangan            string              `json:"keterangan"`
  Nominal               int                 `json:"nominal"`
}

type AddPembayaran struct {
  Id_renter             int                 `json:"id_renter"`
  Id_member             int                 `json:"id_member"`
  Kamar                 string              `json:"kamar"`
  Tipe_pembayaran       string              `json:"tipe_pembayaran"`
  Durasi                string              `json:"durasi"`
  Tanggal_masuk         string              `json:"tanggal_masuk"`
  Tanggal_akhir         string              `json:"tanggal_akhir"`
  Tanggal_penagihan     string              `json:"tanggal_penagihan"`
  Denda                 int                 `json:"denda"`
  Jatuh_tempo           string              `json:"jatuh_tempo"`
  Harga_sewa            int                 `json:"harga_sewa"`
  Total_pembayaran      int                 `json:"total_pembayaran"`
  Total_dibayar         int                 `json:"total_dibayar"`
  Tagihan               int                 `json:"tagihan"`
  Status_pembayaran     string              `json:"status_pembayaran"`
  Tanggal_pembayaran    string              `json:"tanggal_pembayaran"`
}

type AddTanggalPembayaran struct {
  Id_pembayaran         int                 `json:"id_pembayaran"`
  Tanggal_pembayaran    string              `json:"tanggal_pembayaran"`
  Nominal               int                 `json:"nominal"`
}
