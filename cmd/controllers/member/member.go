package member

import (
  "net/http"
  "fmt"
  "crypto/sha1"
  "encoding/json"

  "github.com/labstack/echo"

  m_member "koskuy-ws/cmd/models/member"
)

type M map[string]interface{}

func GetMember(c echo.Context) error {
  id     := c.Param("id")
  member, err := m_member.GetMember(id)

  if err != nil {
    return c.JSON(http.StatusOK, M{"status": "Unauthorized"})
  } else {
    return c.JSON(http.StatusOK, member)
  }
}

func ChangePassword(c echo.Context) error {
  decoder := json.NewDecoder(c.Request().Body)
  data    := struct {
    Id_member            string        `json:"id_member"`
    Password_lama        string        `json:"password_lama"`
    Password_baru        string        `json:"password_baru"`
  }{}
    if err := decoder.Decode(&data); err != nil {
      http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
    }

  var oldPass = sha1.New()
  oldPass.Write([]byte(data.Password_lama))
  var encryptedOldPass = fmt.Sprintf("%x", oldPass.Sum(nil))

  isValid := m_member.CheckOldPassword(data.Id_member, encryptedOldPass)
  if isValid == true {
    var newPass = sha1.New()
    newPass.Write([]byte(data.Password_baru))
    var encryptedNewPass = fmt.Sprintf("%x", newPass.Sum(nil))

    updatePass := m_member.UpdatePassword(data.Id_member,encryptedNewPass)
    return c.JSON(http.StatusOK,  M{"status": updatePass})
  } else {
    return c.JSON(http.StatusOK,  M{"status": false}) // Password lama tidak sesuai.
  }
  return c.NoContent(http.StatusNoContent)
}
