package controllers

import (
	"apis_the_blog/models"
	"encoding/json"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
)

// CategoriasController operations for Categorias
type CategoriasController struct {
	beego.Controller
}

// URLMapping ...
func (c *CategoriasController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Categorias
// @Param	body		body 	models.Categorias	true		"body for Categorias content"
// @Success 201 {int} models.Categorias
// @Failure 403 body is empty
// @router / [post]
func (c *CategoriasController) Post() {
	var v models.Categorias
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddCategorias(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"success": true, "status": 201, "message": "Registro creado exitosamente", "data": v}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": err.Error()}
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": err.Error()}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Categorias by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Categorias
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CategoriasController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCategoriasById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 404, "message": "Registro no encontrado"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición exitosa", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Categorias
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Categorias
// @Failure 403
// @router / [get]
func (c *CategoriasController) GetAll() {
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
				c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error: invalid query key/value pair"}
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllCategorias(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor GetAll: la solicitud contiene un parámetro incorrecto o no contiene registro solicitado"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición exitosa", "data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Categorias
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Categorias	true		"body for Categorias content"
// @Success 200 {object} models.Categorias
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CategoriasController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Categorias{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateCategoriasById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Registro actualizado exitosamente", "data": v}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": err.Error()}
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": err.Error()}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Categorias
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CategoriasController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteCategorias(id); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Registro eliminado exitosamente"}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 404, "message": "Registro no encontrado"}
	}
	c.ServeJSON()
}