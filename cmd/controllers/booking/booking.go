package controllers

import (
  "net/http"
  "encoding/json"
  m_booking "koskuy-ws/cmd/models/booking"
  "github.com/labstack/echo"
)

type M map[string]interface{}

func GetBookingList(c echo.Context) error {
  id_member  := c.Param("id")
  data       := m_booking.GetBookingList(id_member)
  return c.JSON(http.StatusOK, data)
}

func UpdateStatusBooking(c echo.Context) error {
  id_booking := c.Param("id")
  decoder    := json.NewDecoder(c.Request().Body)

  data       := struct {
    Status_booking        string        `json:"status_booking"`
  }{}
    if err := decoder.Decode(&data); err != nil {
      http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
    }

  update_status := m_booking.UpdateStatusBooking(id_booking, data.Status_booking)

  return c.JSON(http.StatusOK, M{"status": update_status})

}
