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

  //  MEMBER
  e.GET("/member/:id", c_member.GetMember, middleware.JWTWithConfig(c_auth.Config))
  e.PUT("/password", c_member.ChangePassword, middleware.JWTWithConfig(c_auth.Config))

  // OWNER
  e.GET("/mykoslist/:id_member", c_kos.GetMykosList, middleware.JWTWithConfig(c_auth.Config))
  e.GET("/mykos", c_kos.GetMykos, middleware.JWTWithConfig(c_auth.Config))
  e.GET("/laporan-pembayaran", c_pembayaran.GetLaporanPembayaran, middleware.JWTWithConfig(c_auth.Config))
  e.GET("/renters", c_renter.GetDaftarRenter, middleware.JWTWithConfig(c_auth.Config))

  fmt.Println("service main started at :8000")
  e.Logger.Fatal(e.Start(":8000"))
}
