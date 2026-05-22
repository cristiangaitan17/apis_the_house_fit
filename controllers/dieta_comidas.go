package controllers

import (
    "encoding/json"
    "strconv"
    "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
    "apis_the_house_fit/models"
)

type DietaComidasController struct {
    web.Controller
}

func (c *DietaComidasController) Get() {
    o := orm.NewOrm()
    var items []*models.DietaComidas
    o.QueryTable(new(models.DietaComidas)).All(&items)
    c.Data["json"] = items
    c.ServeJSON()
}

func (c *DietaComidasController) GetOne() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    o := orm.NewOrm()
    item := models.DietaComidas{Id: id}
    o.Read(&item)
    c.Data["json"] = item
    c.ServeJSON()
}

func (c *DietaComidasController) Post() {
    var m models.DietaComidas
    json.Unmarshal(c.Ctx.Input.RequestBody, &m)
    o := orm.NewOrm()
    o.Insert(&m)
    c.Data["json"] = m
    c.ServeJSON()
}

func (c *DietaComidasController) Put() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    var m models.DietaComidas
    json.Unmarshal(c.Ctx.Input.RequestBody, &m)
    m.Id = id
    o := orm.NewOrm()
    o.Update(&m)
    c.Data["json"] = m
    c.ServeJSON()
}

func (c *DietaComidasController) Delete() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    o := orm.NewOrm()
    o.Delete(&models.DietaComidas{Id: id})
    c.Data["json"] = map[string]string{"message": "Eliminado"}
    c.ServeJSON()
}