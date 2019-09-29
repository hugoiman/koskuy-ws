package structs

import(
  "time"
)

type KosList struct {
  KosList          []Kos               `json:"kos_list"`
}

type Kos struct {
  Id_member        int                 `json:"id_member"`   //  Wewenangan
  Otoritas         string              `json:"otoritas"`
  Status           string              `json:"status"`

  Id_kos           int                 `json:"id_kos"`      //  Kos
  Nama_kos         string              `json:"nama_kos"`
  Tipe_kos         string              `json:"tipe_kos"`
  Alamat           string              `json:"alamat"`
  Luas_kamar       string              `json:"luas_kamar"`
  Total_kamar      int                 `json:"total_kamar"`
  Kamar_terisi     int                 `json:"kamar_terisi"`
  Deskripsi        string              `json:"deskripsi"`
  Verifikasi_kos   bool                `json:"verifikasi_kos"`
  Update_at_ori    time.Time           `json:"update_at_ori"`

  Update_at        string              `json:"update_at"`

  HargaSewaList    []Harga_sewa        `json:"harga_sewa_list"`
  FotoKosList      []Foto_kos          `json:"foto_kos_list"`
  FasilitasKosList []Fasilitas_kos     `json:"fasilitas_kos_list"`

  Review           []Review            `json:"review"`
}

type Foto_kos struct {
  Id_foto_kos      int                 `json:"id_foto_kos"`
  Nama_foto_kos    string              `json:"nama_foto_kos"`
}

type Harga_sewa struct {
  Bulanan          int                 `json:"bulanan"`
  Harian           int                 `json:"harian"`
  Mingguan         int                 `json:"mingguan"`
  Tahunan          int                 `json:"tahunan"`
}

type Fasilitas_kos struct {
  Id_fasilitas     int                 `json:"id_fasilitas"`
  Jenis_fasilitas  string              `json:"jenis_fasilitas"`
  Nama_fasilitas   string              `json:"nama_fasilitas"`
}

type Review struct {
  Ulasan            string              `json:"ulasan"`
  Kebersihan        float64             `json:"kebersihan"`
  Kenyamanan        float64             `json:"kenyamanan"`
  Keamanan          float64             `json:"keamanan"`
  Fasilitas_kamar   float64             `json:"fasilitas_kamar"`
  Fasilitas_bersama float64             `json:"fasilitas_bersama"`
  Harga             float64             `json:"harga"`
  Rating            float64             `json:"rating"`
}
