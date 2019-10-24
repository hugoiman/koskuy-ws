package main

import (
  "fmt"
  "koskuy-ws/cmd/controllers"

  // c_auth "koskuy-ws/cmd/controllers/auth"
  // c_member "koskuy-ws/cmd/controllers/member"
  // c_kos "koskuy-ws/cmd/controllers/kos"
  // c_pembayaran "koskuy-ws/cmd/controllers/pembayaran"
  // c_renter "koskuy-ws/cmd/controllers/renter"
  // c_booking "koskuy-ws/cmd/controllers/booking"
  // c_favorit "koskuy-ws/cmd/controllers/favorit"
  // c_notifikasi "koskuy-ws/cmd/controllers/notifikasi"
  // c_fasilitas "koskuy-ws/cmd/controllers/fasilitas"

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
  e.POST("/authToken", controllers.AuthToken, middleware.JWTWithConfig(controllers.Config))
  e.POST("/login", controllers.Login)
  e.GET("/logout", controllers.Logout)
  e.POST("/member", controllers.RegistrasiMember)
  e.POST("/checkUsername", controllers.CheckUniqueUsername)
  e.POST("/checkEmail", controllers.CheckUniqueEmail)
  e.POST("/forgot-password", controllers.ForgotPassword)
  e.POST("/reset-password", controllers.ResetPassword, middleware.JWTWithConfig(controllers.ConfigResetPassword))

  //  MEMBER
  e.GET("/member/:id", c_member.GetMember, middleware.JWTWithConfig(controllers.Config))
  e.PUT("/member/:id", c_member.EditMember, middleware.JWTWithConfig(controllers.Config))
  e.PUT("/password/:id_member", c_member.ChangePassword, middleware.JWTWithConfig(controllers.Config))
  e.GET("/favorit/:id_member", c_favorit.GetKosFavorit, middleware.JWTWithConfig(controllers.Config))

  // DATA KOS
  e.GET("/mykos/:id_member", c_kos.GetMykosList, middleware.JWTWithConfig(controllers.Config))
  e.GET("/mykos/:id_kos/:id_member", c_kos.GetMykos, middleware.JWTWithConfig(controllers.Config))

  // PEMBAYARAN
  e.GET("/laporan-pembayaran/:id_kos", c_pembayaran.GetLaporanPembayaran, middleware.JWTWithConfig(controllers.Config))
  e.GET("/laporan-bulanan/:id_kos", c_pembayaran.GetLaporanBulanan, middleware.JWTWithConfig(controllers.Config))
  e.GET("/status-pembayaran/:id_kos", c_pembayaran.GetStatusPembayaran, middleware.JWTWithConfig(controllers.Config)) // belum (view)
  e.GET("/pembayaran/:id_pembayaran", c_pembayaran.GetPembayaran, middleware.JWTWithConfig(controllers.Config)) // belum (view)
  e.POST("/pembayaran", c_pembayaran.AddPembayaran, middleware.JWTWithConfig(controllers.Config))
  e.GET("/histori-pembayaran/:id_member", c_pembayaran.GetHistoryPembayaran, middleware.JWTWithConfig(controllers.Config))



  // RENTER
  e.GET("/daftar-anak-kos/:id_kos", c_renter.GetDaftarRenter, middleware.JWTWithConfig(controllers.Config))  // belum (view)
  e.GET("/anak-kos/:slug", c_renter.GetRenter, middleware.JWTWithConfig(controllers.Config))    // belum  (view)

  //  BOOKING
  e.GET("/booking/:id", c_booking.GetBookingList, middleware.JWTWithConfig(controllers.Config))  // id_member
  e.PUT("/booking/:id", c_booking.UpdateStatusBooking, middleware.JWTWithConfig(controllers.Config)) //  id_booking

  //  Notifikasi
  e.GET("/notifikasi_booking/:id_member", c_notifikasi.GetNotifikasiBookingList, middleware.JWTWithConfig(controllers.Config))  // id_member
  // e.GET("/total_notifikasi_booking/:id_member", c_notifikasi.GetNotifikasiList, middleware.JWTWithConfig(controllers.Config))  // id_member

  // Fasilitas
  e.GET("/all-fasilitas", c_fasilitas.GetAllFasilitas, middleware.JWTWithConfig(controllers.Config))

  fmt.Println("service main started at :8000")
  e.Logger.Fatal(e.Start(":8000"))
}
