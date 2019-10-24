package main

import (
  "fmt"
  "koskuy-ws/cmd/controllers"

  // c_auth "koskuy-ws/cmd/controllers/auth"
  // controllers "koskuy-ws/cmd/controllers/member"
  // controllers "koskuy-ws/cmd/controllers/kos"
  // controllers "koskuy-ws/cmd/controllers/pembayaran"
  // controllers "koskuy-ws/cmd/controllers/renter"
  // controllers "koskuy-ws/cmd/controllers/booking"
  // c_favorit "koskuy-ws/cmd/controllers/favorit"
  // controllers "koskuy-ws/cmd/controllers/notifikasi"
  // controllers "koskuy-ws/cmd/controllers/fasilitas"

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
  e.GET("/member/:id", controllers.GetMember, middleware.JWTWithConfig(controllers.Config))
  e.PUT("/member/:id", controllers.EditMember, middleware.JWTWithConfig(controllers.Config))
  e.PUT("/password/:id_member", controllers.ChangePassword, middleware.JWTWithConfig(controllers.Config))
  e.GET("/favorit/:id_member", controllers.GetKosFavorit, middleware.JWTWithConfig(controllers.Config))

  // DATA KOS
  e.GET("/mykos/:id_member", controllers.GetMykosList, middleware.JWTWithConfig(controllers.Config))
  e.GET("/mykos/:id_kos/:id_member", controllers.GetMykos, middleware.JWTWithConfig(controllers.Config))

  // PEMBAYARAN
  e.GET("/laporan-pembayaran/:id_kos", controllers.GetLaporanPembayaran, middleware.JWTWithConfig(controllers.Config))
  e.GET("/laporan-bulanan/:id_kos", controllers.GetLaporanBulanan, middleware.JWTWithConfig(controllers.Config))
  e.GET("/status-pembayaran/:id_kos", controllers.GetStatusPembayaran, middleware.JWTWithConfig(controllers.Config)) // belum (view)
  e.GET("/pembayaran/:id_pembayaran", controllers.GetPembayaran, middleware.JWTWithConfig(controllers.Config)) // belum (view)
  e.POST("/pembayaran", controllers.AddPembayaran, middleware.JWTWithConfig(controllers.Config))
  e.GET("/history-pembayaran/:id_renter/:id_kos", controllers.GetHistoryPembayaranRenter, middleware.JWTWithConfig(controllers.Config))



  // RENTER
  e.GET("/daftar-anak-kos/:id_kos", controllers.GetDaftarRenter, middleware.JWTWithConfig(controllers.Config))  // belum (view)
  e.GET("/anak-kos/:slug", controllers.GetRenter, middleware.JWTWithConfig(controllers.Config))    // belum  (view)

  //  BOOKING
  e.GET("/booking/:id", controllers.GetBookingList, middleware.JWTWithConfig(controllers.Config))  // id_member
  e.PUT("/booking/:id", controllers.UpdateStatusBooking, middleware.JWTWithConfig(controllers.Config)) //  id_booking

  //  Notifikasi
  e.GET("/notifikasi_booking/:id_member", controllers.GetNotifikasiBookingList, middleware.JWTWithConfig(controllers.Config))  // id_member
  // e.GET("/total_notifikasi_booking/:id_member", controllers.GetNotifikasiList, middleware.JWTWithConfig(controllers.Config))  // id_member

  // Fasilitas
  e.GET("/all-fasilitas", controllers.GetAllFasilitas, middleware.JWTWithConfig(controllers.Config))

  fmt.Println("service main started at :8000")
  e.Logger.Fatal(e.Start(":8000"))
}
