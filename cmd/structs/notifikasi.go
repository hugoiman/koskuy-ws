package structs

import(
  "time"
)

type NotifikasiList struct {
  NotifikasiList       []Notifikasi        `json:"notifikasi_list"`
}

type Notifikasi struct {
  Id_notifikasi        int                 `json:"id_notifikasi"`
  Id_member            int                 `json:"id_member"`
  Subjek               string              `json:"subjek"`
  Pesan                string              `json:"pesan"`
  Status_notif         int                 `json:"status_notif"`
  Date_ori             time.Time           `json:"date_ori"`

  Date                 string              `json:"date"`
}
