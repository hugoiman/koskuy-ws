package structs

type AllFasilitas struct {
  AllFasilitas      []Fasilitas        `json:"all_fasilitas"`
}

type Fasilitas struct {
  Id_fasilitas     string              `json:"id_fasilitas"`
  Jenis_fasilitas  string              `json:"jenis_fasilitas"`
  Nama_fasilitas   string              `json:"nama_fasilitas"`
}
