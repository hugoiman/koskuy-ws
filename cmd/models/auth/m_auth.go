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

func CheckUniqueUsername(username string, id_member int) bool {
	var isUnique string
  con     :=  db.Connect()
  query   :=  "SELECT username FROM member WHERE username = ? AND id_member != ?"
  err     :=  con.QueryRow(query, username, username).Scan(&isUnique)

  defer con.Close()

  if err == nil {
    return false
  } else {
    return true
  }
}

func CheckUniqueEmail(email string, id_member int) bool {
	var isUnique string
  con     :=  db.Connect()
  query   :=  "SELECT email FROM member WHERE email = ? AND id_member != ?"
  err     :=  con.QueryRow(query, email, id_member).Scan(&isUnique)
  defer con.Close()

  if err == nil {
    return false
  } else {
    return true
  }
}

func CreateMember(nama, username, email, password string) bool {
  foto    := "member_wevmtv.png"
  con     :=  db.Connect()
	_, err 	:=  con.Exec("INSERT INTO member (nama, username, email, password, foto) VALUES (?,?,?,?,?)", nama, username, email, password, foto)

	defer con.Close()

  if err == nil {
    return true
  } else {
    return false
  }
}
