package controllers

import (
    "encoding/json"
    "strconv"
    "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
    "apis_the_house_fit/models"
)

type CategoriasController struct {
    web.Controller
}

// GET /api/v1/categorias
func (c *CategoriasController) Get() {
    o := orm.NewOrm()
    var items []*models.Categorias
    _, err := o.QueryTable(new(models.Categorias)).All(&items)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
    } else {
        c.Data["json"] = items
    }
    c.ServeJSON()
}

// GET /api/v1/categorias/:id
func (c *CategoriasController) GetOne() {
    id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
    if err != nil {
        c.Data["json"] = map[string]string{"error": "ID inválido"}
        c.ServeJSON()
        return
    }
    
    o := orm.NewOrm()
    item := models.Categorias{Id: id}
    err = o.Read(&item)
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Categoría no encontrada"}
    } else {
        c.Data["json"] = item
    }
    c.ServeJSON()
}

// POST /api/v1/categorias
func (c *CategoriasController) Post() {
    var m models.Categorias
    err := json.Unmarshal(c.Ctx.Input.RequestBody, &m)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }
    
    o := orm.NewOrm()
    id, err := o.Insert(&m)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
    } else {
        m.Id = int(id)
        c.Data["json"] = m
    }
    c.ServeJSON()
}

// PUT /api/v1/categorias/:id
func (c *CategoriasController) Put() {
    id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
    if err != nil {
        c.Data["json"] = map[string]string{"error": "ID inválido"}
        c.ServeJSON()
        return
    }

    var m models.Categorias
    err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

    m.Id = id
    o := orm.NewOrm()
    _, err = o.Update(&m)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
    } else {
        c.Data["json"] = m
    }
    c.ServeJSON()
}

// DELETE /api/v1/categorias/:id
func (c *CategoriasController) Delete() {
    id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
    if err != nil {
        c.Data["json"] = map[string]string{"error": "ID inválido"}
        c.ServeJSON()
        return
    }

    o := orm.NewOrm()
    item := models.Categorias{Id: id}
    _, err = o.Delete(&item)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
    } else {
        c.Data["json"] = map[string]string{"message": "Categoría eliminada"}
    }
    c.ServeJSON()
}