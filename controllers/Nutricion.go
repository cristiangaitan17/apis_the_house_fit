package controllers

import (
    "encoding/json"
    "strconv"
    "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
    "apis_the_house_fit/models"
)

type NutricionController struct {
    web.Controller
}

func (c *NutricionController) Get() {
    o := orm.NewOrm()
    var items []*models.Nutricion
    o.QueryTable(new(models.Nutricion)).All(&items)
    c.Data["json"] = items
    c.ServeJSON()
}

func (c *NutricionController) GetOne() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    o := orm.NewOrm()
    item := models.Nutricion{Id: id}
    o.Read(&item)
    c.Data["json"] = item
    c.ServeJSON()
}

func (c *NutricionController) Post() {
    var m models.Nutricion
    json.Unmarshal(c.Ctx.Input.RequestBody, &m)
    o := orm.NewOrm()
    o.Insert(&m)
    c.Data["json"] = m
    c.ServeJSON()
}

func (c *NutricionController) Put() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    var m models.Nutricion
    json.Unmarshal(c.Ctx.Input.RequestBody, &m)
    m.Id = id
    o := orm.NewOrm()
    o.Update(&m)
    c.Data["json"] = m
    c.ServeJSON()
}

func (c *NutricionController) Delete() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    o := orm.NewOrm()
    o.Delete(&models.Nutricion{Id: id})
    c.Data["json"] = map[string]string{"message": "Eliminado"}
    c.ServeJSON()
}