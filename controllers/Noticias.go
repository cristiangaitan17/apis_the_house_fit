package controllers

import (
    "encoding/json"
    "strconv"
    "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
    "apis_the_house_fit/models"
)

type NoticiasController struct {
    web.Controller
}

func (c *NoticiasController) Get() {
    o := orm.NewOrm()
    var items []*models.Noticias
    o.QueryTable(new(models.Noticias)).All(&items)
    c.Data["json"] = items
    c.ServeJSON()
}

func (c *NoticiasController) GetOne() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    o := orm.NewOrm()
    item := models.Noticias{Id: id}
    o.Read(&item)
    c.Data["json"] = item
    c.ServeJSON()
}

func (c *NoticiasController) Post() {
    var m models.Noticias
    json.Unmarshal(c.Ctx.Input.RequestBody, &m)
    o := orm.NewOrm()
    o.Insert(&m)
    c.Data["json"] = m
    c.ServeJSON()
}

func (c *NoticiasController) Put() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    var m models.Noticias
    json.Unmarshal(c.Ctx.Input.RequestBody, &m)
    m.Id = id
    o := orm.NewOrm()
    o.Update(&m)
    c.Data["json"] = m
    c.ServeJSON()
}

func (c *NoticiasController) Delete() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    o := orm.NewOrm()
    o.Delete(&models.Noticias{Id: id})
    c.Data["json"] = map[string]string{"message": "Eliminado"}
    c.ServeJSON()
}