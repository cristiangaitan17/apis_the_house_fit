package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"the_house_fit/models"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// AdministradorController operations for Administrador
type AdministradorController struct {
	beego.Controller
}

// URLMapping ...
func (c *AdministradorController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @router / [post]
func (c *AdministradorController) Post() {
	var v models.Administrador
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddAdministrador(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"success": true, "status": 201, "Message": "Administrador creado exitosamente", "data": v}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error al crear el Administrador: " + err.Error(), "data": nil}
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error al leer el cuerpo de la solicitud: " + err.Error(), "data": nil}
	}
	c.ServeJSON()
}

// GetOne ...
// @router /:id [get]
func (c *AdministradorController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAdministradorById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servicio GetOne: La solicitud contiene un parametro incorrecto o no existe ningun registro", "data": nil}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @router / [get]
func (c *AdministradorController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error: par clave/valor de query invalido", "data": nil}
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllAdministrador(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor GetAll: la solicitud contiene un parametro incorrecto o no contiene registro solicitado", "data": nil}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa", "data": l}
	}
	c.ServeJSON()
}

// Put ...
// @router /:id [put]
func (c *AdministradorController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Administrador{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateAdministradorById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Administrador actualizado exitosamente", "data": v}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error al actualizar el Administrador: " + err.Error(), "data": nil}
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error al leer el cuerpo de la solicitud: " + err.Error(), "data": nil}
	}
	c.ServeJSON()
}

// Delete ...
// @router /:id [delete]
func (c *AdministradorController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAdministrador(id); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Administrador eliminado exitosamente", "data": nil}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error al eliminar el Administrador: " + err.Error(), "data": nil}
	}
	c.ServeJSON()
}
