package auth_member

import(
  "fmt"
  "koskuy-ws/db"
)

func IsAuth(id, password string) bool {
	var isAuth string
	con     :=  db.Connect()
  query   :=  "SELECT email FROM member WHERE (email = ? OR no_hp = ? OR username = ?) AND password = ?"
  err     :=  con.QueryRow(query, id, id, id, password).Scan(&isAuth)

	defer con.Close()

	if err == nil {
		return true
	} else {
		return false
	}
}

func GetIdMember(id string) string {
  var id_member string
  con     :=  db.Connect()
  query   :=  "SELECT id_member FROM member WHERE email = ? OR no_hp = ? OR username = ?"
  err     :=  con.QueryRow(query, id, id, id).Scan(&id_member)

  if err != nil {
    fmt.Println(err.Error())
  }

  defer con.Close()

  return id_member
}
