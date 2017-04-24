package controllers

import (
  "cdve_api_crud/models"
  "github.com/astaxie/beego"
)

type ContratadoController struct {
	beego.Controller
}

func (c *ContratadoController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)

}


func (c *ContratadoController) GetAll() {
    listaContratados := models.GetAllContratado()
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaContratados
    c.ServeJSON()
}