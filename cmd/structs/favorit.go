package structs

import(
  "time"
)

type FavoritList struct {
  FavoritList      []Favorit           `json:"favorit_list"`
}

type Favorit struct {
  Id_kos           int                 `json:"id_kos"`
  Id_member        int                 `json:"id_member"`

  Nama_kos         string              `json:"nama_kos"`
  Tipe_kos         string              `json:"tipe_kos"`
  Booking          string              `json:"booking"`
  Slug             string              `json:"slug"`
  Update_at_ori    time.Time           `json:"update_at_ori"`

  Update_at        string              `json:"update_at"`

  HargaSewaList    []Harga_sewa        `json:"harga_sewa_list"`
  FotoKosList      []Foto_kos          `json:"foto_kos_list"`
}
