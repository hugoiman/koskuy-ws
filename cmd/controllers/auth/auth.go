package auth

import (
  "net/http"
  "fmt"
  "time"
  "crypto/sha1"
  "encoding/json"

  "koskuy-ws/cmd/structs"

  c_email "koskuy-ws/cmd/controllers/email"
  m_auth "koskuy-ws/cmd/models/auth"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/dgrijalva/jwt-go"
)

const COOKIE_NAME = "cookie_token"
var expirationTime = time.Now().Add(time.Hour * 60)
var expirationTimeVerification = time.Now().Add(time.Minute * 30)
var mySigningKey = []byte("mysupersecret")

var Config = middleware.JWTConfig{
  Claims:     &structs.JwtCustomClaims{},
  SigningKey: mySigningKey,
}

var ConfigResetPassword = middleware.JWTConfig{
  Claims:     &structs.JwtResetPassword{},
  SigningKey: mySigningKey,
}

type M map[string]interface{}

func Login(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
  data    := struct {
    Id            string        `json:"id"`
    Password      string        `json:"password"`
  }{}
    if err := decoder.Decode(&data); err != nil {
      http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
    }

  var sha = sha1.New()
  sha.Write([]byte(data.Password))
  var encrypted = sha.Sum(nil)
  var encryptedString = fmt.Sprintf("%x", encrypted)

  authentic    := m_auth.IsAuth(data.Id, encryptedString)

  if authentic == true {
    id_member   := m_auth.GetIdMember(data.Id)
    tokenString := GenerateJWT(c, id_member, "member")
    return c.JSON(http.StatusOK, M{"token": tokenString})
	} else {
    return c.JSON(http.StatusOK,M{"status": false})
	}
}

func GenerateJWT(c echo.Context, id_user, user string) string {
  // Set custom claims
  claims := &structs.JwtCustomClaims{
		Id_user: id_user,
    User: user,
    StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //  Create token with claim
  tokenString, err := token.SignedString(mySigningKey)      // Generate encoded token and send it as response.
  if err != nil {
    return "Generate token error"
  }

  return tokenString
}

func AuthToken(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(*structs.JwtCustomClaims)
  data := struct {
      Id_user   string
      User      string
    }{
      claims.Id_user,
      claims.User,
    }
	return c.JSON(http.StatusOK, data)
}

func Logout(c echo.Context) error  {
  cookie := &http.Cookie{
    Name:     COOKIE_NAME,
    Value:    "",
    Path:     "/",
    Expires: time.Unix(0, 0),
  }

  http.SetCookie(c.Response(), cookie)
  return c.Redirect(http.StatusMovedPermanently, "http://localhost:9000/auth")
}

func RegistrasiMember(c echo.Context) error  {
  decoder := json.NewDecoder(c.Request().Body)
  data    := struct {
    Nama           string        `json:"nama"`
    Username       string        `json:"username"`
    Email          string        `json:"email"`
    Password       string        `json:"password"`
  }{}
    if err := decoder.Decode(&data); err != nil {
      http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
    }

  var pass = sha1.New()
  pass.Write([]byte(data.Password))
  var encryptedPassword = fmt.Sprintf("%x", pass.Sum(nil))

  create_member := m_auth.CreateMember(data.Nama, data.Username, data.Email, encryptedPassword)

  return c.JSON(http.StatusOK, M{"status": create_member})
}

func CheckUniqueUsername(c echo.Context) error {
  decoder := json.NewDecoder(c.Request().Body)
  data    := struct {
    Id_member           int           `json:"id_member"`
    Username            string        `json:"username"`
  }{}
    if err := decoder.Decode(&data); err != nil {
      http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
    }

  isUnique := m_auth.CheckUniqueUsername(data.Username, data.Id_member)
  return c.JSON(http.StatusOK, M{"status": isUnique})
}

func CheckUniqueEmail(c echo.Context) error {
  decoder := json.NewDecoder(c.Request().Body)
  data    := struct {
    Id_member           int           `json:"id_member"`
    Email               string        `json:"email"`
  }{}
    if err := decoder.Decode(&data); err != nil {
      http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
    }

  isUnique := m_auth.CheckUniqueEmail(data.Email, data.Id_member)
  return c.JSON(http.StatusOK, M{"status": isUnique})
}

func ForgotPassword(c echo.Context) error {
  decoder := json.NewDecoder(c.Request().Body)
  data    := struct {
    Email        string        `json:"email"`
  }{}
    if err := decoder.Decode(&data); err != nil {
      http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
    }

  token     := GenerateJWTResetPassword(c, data.Email)
  link      := "http://localhost:9000/reset-password?token="+token

  subjek    := "Koskuy - [Reset Password]"
  pesan     := "Hallo,<br<br>Email ini telah dikirimkan agar kamu dapat mengatur ulang password.<br><br>Silahkan klik link dibawah ini untuk me-reset password kamu. Link ini akan kadaluarsa dalam waktu 30 menit.<br><br>"+link

  sent      := c_email.SendEmail(data.Email, subjek, pesan)
  return c.JSON(http.StatusOK, M{"status": sent}) // true/false
}

func GenerateJWTResetPassword(c echo.Context, email string) string {
  // Set custom claims
  claims := &structs.JwtResetPassword{
		Email: email,
    StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTimeVerification.Unix(),
		},
	}

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //  Create token with claim
  tokenString, _ := token.SignedString(mySigningKey)      // Generate encoded token and send it as response.
  // if err != nil {
  //   return "Generate token error"
  // }

  return tokenString
}

func ResetPassword(c echo.Context) error  {
  decoder := json.NewDecoder(c.Request().Body)
  data    := struct {
    Password        string        `json:"password"`
  }{}
    if err := decoder.Decode(&data); err != nil {
      http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
    }

  // Decode token
  user   := c.Get("user").(*jwt.Token)
  claims := user.Claims.(*structs.JwtResetPassword)
  data_token := struct {
      Email   string
    }{
      claims.Email,
    }
  //  Enkripsi Password
  var pass = sha1.New()
  pass.Write([]byte(data.Password))
  var encryptedPassword = fmt.Sprintf("%x", pass.Sum(nil))

  //  Update Password
  update_password := m_auth.UpdatePassword(data_token.Email, encryptedPassword)

  return c.JSON(http.StatusOK, M{"status": update_password}) // true/false
}
