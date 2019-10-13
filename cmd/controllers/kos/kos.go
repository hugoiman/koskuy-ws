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
  id        := c.Param("id_kos")  // slug/id_kos
  id_member := c.Param("id_member")
  data,_    := m_kos.GetMyKos(id, id_member)
  return c.JSON(http.StatusOK, data)
}
