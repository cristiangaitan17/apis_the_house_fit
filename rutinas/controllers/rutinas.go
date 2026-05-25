package controllers

import (
	"encoding/json"
	"errors"
	"rutinas/models"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// RutinasController operations for Rutinas
type RutinasController struct {
	beego.Controller
}

// URLMapping ...
func (c *RutinasController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Rutinas
// @Param	body		body 	models.Rutinas	true		"body for Rutinas content"
// @Success 201 {int} models.Rutinas
// @Failure 403 body is empty
// @router / [post]
func (c *RutinasController) Post() {
	var v models.Rutinas
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor: la solicitud contiene un parametro incorrecto"}
	} else {
		if _, err := models.AddRutinas(&v); err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error en el servidor: no se pudo crear la rutina"}
		} else {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"success": true, "status": 201, "message": "Rutina creada", "data": v}
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Rutinas by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Rutinas
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RutinasController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetRutinasById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 404, "message": "Error en el servidor: rutina no encontrada"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Rutina obtenida", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Rutinas
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Rutinas
// @Failure 403
// @router / [get]
func (c *RutinasController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllRutinas(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error en el servidor: no se pudo obtener las rutinas"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Rutinas obtenidas", "data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Rutinas
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Rutinas	true		"body for Rutinas content"
// @Success 200 {object} models.Rutinas
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RutinasController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Rutinas{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor: la solicitud contiene un parametro incorrecto"}
	} else {
		if err := models.UpdateRutinasById(&v); err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error en el servidor: no se pudo actualizar la rutina"}
		} else {
			c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Rutina actualizada", "data": v}
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Rutinas
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RutinasController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteRutinas(id); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error en el servidor: no se pudo eliminar la rutina"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Rutina eliminada"}
	}
	c.ServeJSON()
}
