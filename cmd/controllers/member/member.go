package member

import (
  "net/http"
  "fmt"
  "crypto/sha1"
  "encoding/json"

  "koskuy-ws/cmd/structs"

  m_member "koskuy-ws/cmd/models/member"

  "github.com/labstack/echo"
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
  id_member := c.Param("id_member")
  decoder := json.NewDecoder(c.Request().Body)
  data    := struct {
    Password_lama        string        `json:"password_lama"`
    Password_baru        string        `json:"password_baru"`
  }{}
    if err := decoder.Decode(&data); err != nil {
      http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
    }

  var oldPass = sha1.New()
  oldPass.Write([]byte(data.Password_lama))
  var encryptedOldPass = fmt.Sprintf("%x", oldPass.Sum(nil))

  isValid := m_member.CheckOldPassword(id_member, encryptedOldPass)
  if isValid == true {
    var newPass = sha1.New()
    newPass.Write([]byte(data.Password_baru))
    var encryptedNewPass = fmt.Sprintf("%x", newPass.Sum(nil))

    updatePass := m_member.UpdatePassword(id_member,encryptedNewPass)
    return c.JSON(http.StatusOK,  M{"status": updatePass})
  } else {
    return c.JSON(http.StatusOK,  M{"status": false}) // Password lama tidak sesuai.
  }
  return c.NoContent(http.StatusNoContent)
}

func EditMember(c echo.Context) error {
  // id_member := c.Param("id_member")
  decoder := json.NewDecoder(c.Request().Body)
  data    := structs.Member{}
  decoder.Decode(&data)
  update_member := m_member.UpdateMember(data.Id_member, data)
  return c.JSON(http.StatusOK, M{"status": update_member})
}
