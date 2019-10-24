package controllers

import (
  "net/http"
  "koskuy-ws/cmd/models"
  "github.com/labstack/echo"
)

func GetMykosList(c echo.Context) error {
  id_member := c.Param("id_member")
  data      := models.GetMykosList(id_member)
  return c.JSON(http.StatusOK, data)
}

func GetMykos(c echo.Context) error {
  id        := c.Param("id_kos")  // slug/id_kos
  id_member := c.Param("id_member")
  data,_    := models.GetMyKos(id, id_member)
  return c.JSON(http.StatusOK, data)
}
