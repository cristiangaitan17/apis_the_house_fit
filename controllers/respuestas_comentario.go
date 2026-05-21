package controllers

import (
    "encoding/json"
    "strconv"
    "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
    "apis_the_house_fit/models"
)

type RespuestasComentarioController struct {
    web.Controller
}

func (c *RespuestasComentarioController) Get() {
    o := orm.NewOrm()
    var items []*models.RespuestasComentario
    o.QueryTable(new(models.RespuestasComentario)).All(&items)
    c.Data["json"] = items
    c.ServeJSON()
}

func (c *RespuestasComentarioController) GetOne() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    o := orm.NewOrm()
    item := models.RespuestasComentario{Id: id}
    o.Read(&item)
    c.Data["json"] = item
    c.ServeJSON()
}

func (c *RespuestasComentarioController) Post() {
    var m models.RespuestasComentario
    json.Unmarshal(c.Ctx.Input.RequestBody, &m)
    o := orm.NewOrm()
    o.Insert(&m)
    c.Data["json"] = m
    c.ServeJSON()
}

func (c *RespuestasComentarioController) Put() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    var m models.RespuestasComentario
    json.Unmarshal(c.Ctx.Input.RequestBody, &m)
    m.Id = id
    o := orm.NewOrm()
    o.Update(&m)
    c.Data["json"] = m
    c.ServeJSON()
}

func (c *RespuestasComentarioController) Delete() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    o := orm.NewOrm()
    o.Delete(&models.RespuestasComentario{Id: id})
    c.Data["json"] = map[string]string{"message": "Eliminado"}
    c.ServeJSON()
}