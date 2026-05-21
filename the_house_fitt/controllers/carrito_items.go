package controllers

import (
	"the_house_fitt/models"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// CarritoItemsController operations for CarritoItems
type CarritoItemsController struct {
	beego.Controller
}

// URLMapping ...
func (c *CarritoItemsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create CarritoItems
// @Param	body		body 	models.CarritoItems	true		"body for CarritoItems content"
// @Success 201 {int} models.CarritoItems
// @Failure 403 body is empty
// @router / [post]
func (c *CarritoItemsController) Post() {
	var v models.CarritoItems
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		logs.Error(err)
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servicio Post: el cuerpo de la solicitud es inválido o está mal formado"}
		c.ServeJSON()
		return
	}
		
		
	if _, err := models.AddCarritoItems(&v); err != nil {
		logs.Error(err)
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servicio Post: no se pudo crear el CarritoItems a la base de datos"}
		c.ServeJSON()
		return
	}
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{"success": true, "status": 201, "message": "CarritoItems creado exitosamente", "data": v}
	c.ServeJSON()
			
}

// GetOne ...
// @Title Get One
// @Description get CarritoItems by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.CarritoItems
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CarritoItemsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servicio GetOne: el id debe ser un número entero válido"}
		c.ServeJSON()
		return
	}

	v, err := models.GetCarritoItemsById(id)
	if err != nil {
		logs.Error(err)
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 404, "message": "Error en el servicio GetOne: la solicitud contiene un parámetro incorrecto o no existe el registro"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición exitosa", "data": v}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get CarritoItems
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.CarritoItems
// @Failure 403
// @router / [get]
func (c *CarritoItemsController) GetAll() {
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
				c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servicio GetAll: parametro de consultainválido"}
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllCarritoItems(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error en el servicio GetAll: no se pudieron obtener los registros"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición exitosa", "data": l}
	c.ServeJSON()
}

	

// Put ...
// @Title Put
// @Description update the CarritoItems
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.CarritoItems	true		"body for CarritoItems content"
// @Success 200 {object} models.CarritoItems
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CarritoItemsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servicio Put: el id debe ser un número entero válido"}
		c.ServeJSON()
		return
	}

	v := models.CarritoItems{IdCarrito: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servicio Put: el cuerpo de la solicitud es inválido o está mal formado"}
		c.ServeJSON()
		return
	}

	if err := models.UpdateCarritoItemsById(&v); err != nil {
		logs.Error(err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error en el servicio Put: no se pudo actualizar el registro"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Registro actualizado exitosamente"}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the CarritoItems
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CarritoItemsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servicio Delete: el id debe ser un número entero válido"}
		c.ServeJSON()
		return
	}
	if err := models.DeleteCarritoItems(id); err != nil {
		logs.Error(err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error en el servicio Delete: no se pudo eliminar el registro"}
		c.ServeJSON()
		return
	}
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Registro eliminado exitosamente"}
	c.ServeJSON()
}
