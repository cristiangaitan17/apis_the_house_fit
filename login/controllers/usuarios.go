package controllers

import (
	"encoding/json"
	"strconv"
	"strings"
	"the_house_fit/models"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// UsuariosController operations for Usuarios
type UsuariosController struct {
	beego.Controller
}

// URLMapping ...
func (c *UsuariosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @router / [post]
func (c *UsuariosController) Post() {
	var v models.Usuarios
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddUsuarios(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"success": true, "status": 201, "Message": "Usuario creado exitosamente", "data": v}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error al crear el Usuario: " + err.Error(), "data": nil}
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error al leer el cuerpo de la solicitud: " + err.Error(), "data": nil}
	}
	c.ServeJSON()
}

// GetOne ...
// @router /:id [get]
func (c *UsuariosController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUsuariosById(id)
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
func (c *UsuariosController) GetAll() {
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

	l, err := models.GetAllUsuarios(query, fields, sortby, order, offset, limit)
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
func (c *UsuariosController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Usuarios{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUsuariosById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Usuario actualizado exitosamente", "data": v}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error al actualizar el Usuario: " + err.Error(), "data": nil}
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error al leer el cuerpo de la solicitud: " + err.Error(), "data": nil}
	}
	c.ServeJSON()
}

// Delete ...
// @router /:id [delete]
func (c *UsuariosController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUsuarios(id); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Usuario eliminado exitosamente", "data": nil}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error al eliminar el Usuario: " + err.Error(), "data": nil}
	}
	c.ServeJSON()
}
