package controllers

import (
  "net/http"
  m_kos "koskuy-ws/cmd/models/kos"
  "github.com/labstack/echo"
)

func GetMykosList(c echo.Context) error {
  id_member := c.Param("id_member")
  data      := m_kos.GetMykosList(id_member)
  return c.JSON(http.StatusOK, data)
}

func GetMykos(c echo.Context) error {
  id_kos    := c.QueryParam("kos")
  id_member := c.QueryParam("member")
  data,_    := m_kos.GetMyKos(id_kos, id_member)
  return c.JSON(http.StatusOK, data)
}
