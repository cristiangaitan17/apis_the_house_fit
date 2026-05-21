package controllers

import (
    "encoding/json"
    "strconv"
    "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
    "apis_the_house_fit/models"
)

type ComentariosComunidadController struct {
    web.Controller
}

func (c *ComentariosComunidadController) Get() {
    o := orm.NewOrm()
    var items []*models.ComentariosComunidad
    o.QueryTable(new(models.ComentariosComunidad)).All(&items)
    c.Data["json"] = items
    c.ServeJSON()
}

func (c *ComentariosComunidadController) GetOne() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    o := orm.NewOrm()
    item := models.ComentariosComunidad{Id: id}
    o.Read(&item)
    c.Data["json"] = item
    c.ServeJSON()
}

func (c *ComentariosComunidadController) Post() {
    var m models.ComentariosComunidad
    json.Unmarshal(c.Ctx.Input.RequestBody, &m)
    o := orm.NewOrm()
    o.Insert(&m)
    c.Data["json"] = m
    c.ServeJSON()
}

func (c *ComentariosComunidadController) Put() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    var m models.ComentariosComunidad
    json.Unmarshal(c.Ctx.Input.RequestBody, &m)
    m.Id = id
    o := orm.NewOrm()
    o.Update(&m)
    c.Data["json"] = m
    c.ServeJSON()
}

func (c *ComentariosComunidadController) Delete() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    o := orm.NewOrm()
    o.Delete(&models.ComentariosComunidad{Id: id})
    c.Data["json"] = map[string]string{"message": "Eliminado"}
    c.ServeJSON()
}