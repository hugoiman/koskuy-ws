package structs

type AddTanggalPembayaran struct {
  Id_pembayaran         int                 `json:"id_pembayaran"`
  Tanggal_pembayaran    string              `json:"tanggal_pembayaran"`
  Nominal               int                 `json:"nominal"`
}
