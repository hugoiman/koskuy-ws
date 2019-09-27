package main

import (
  "fmt"

  c_auth "koskuy-ws/cmd/controllers/auth"
  c_member "koskuy-ws/cmd/controllers/member"
  c_kos "koskuy-ws/cmd/controllers/kos"
  c_pembayaran "koskuy-ws/cmd/controllers/pembayaran"
  c_renter "koskuy-ws/cmd/controllers/renter"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/gorilla/context"
  "github.com/rs/cors"
)

func main() {
  e := echo.New()

  e.Use(middleware.Recover())
  e.Use(echo.WrapMiddleware(context.ClearHandler))

  corsMiddleware := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
    AllowedHeaders: []string{"Content-Type", "X-CSRF-Token", "Authorization"},
    AllowCredentials: true,
    // Debug:          true,
  })
  e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

  // Endpoint

  //  Authentication
  e.POST("/authToken", c_auth.AuthToken, middleware.JWTWithConfig(c_auth.Config))
  e.POST("/login", c_auth.Login)
  e.GET("/logout", c_auth.Logout)
  e.POST("/member", c_auth.RegistrasiMember)
  e.POST("/checkUsername", c_auth.CheckUniqueUsername)
  e.POST("/checkEmail", c_auth.CheckUniqueEmail)
  e.POST("/forgot-password", c_auth.ForgotPassword)
  e.POST("/reset-password", c_auth.ResetPassword, middleware.JWTWithConfig(c_auth.ConfigResetPassword))

  //  MEMBER
  e.GET("/member/:id", c_member.GetMember, middleware.JWTWithConfig(c_auth.Config))
  e.PUT("/member/:id_member", c_member.EditMember, middleware.JWTWithConfig(c_auth.Config))
  e.PUT("/password/:id_member", c_member.ChangePassword, middleware.JWTWithConfig(c_auth.Config))

  // DATA KOS
  e.GET("/mykos/:id_member", c_kos.GetMykosList, middleware.JWTWithConfig(c_auth.Config))
  e.GET("/mykos/:id_kos/:id_member", c_kos.GetMykos, middleware.JWTWithConfig(c_auth.Config))

  // PEMBAYARAN
  e.GET("/laporan-pembayaran/:id_kos", c_pembayaran.GetLaporanPembayaran, middleware.JWTWithConfig(c_auth.Config))
  e.GET("/history-pembayaran/:id_renter", c_pembayaran.GetHistoryPembayaran, middleware.JWTWithConfig(c_auth.Config)) //  belum(view, logic)
  e.GET("/pembayaran/:id_pembayaran", c_pembayaran.GetPembayaran, middleware.JWTWithConfig(c_auth.Config)) // belum (view)

  // RENTER
  e.GET("/daftar-anak-kos/:id_kos", c_renter.GetDaftarRenter, middleware.JWTWithConfig(c_auth.Config))  // belum (view)
  e.GET("/anak-kos/:id_renter", c_renter.GetRenter, middleware.JWTWithConfig(c_auth.Config))    // belum  (view)

  fmt.Println("service main started at :8000")
  e.Logger.Fatal(e.Start(":8000"))
}
