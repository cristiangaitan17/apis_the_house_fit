package controllers

import (
	"encoding/json"
	"entrenadores/models"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// SolicitudesEntrenadorController operations for SolicitudesEntrenador
type SolicitudesEntrenadorController struct {
	beego.Controller
}

// URLMapping ...
func (c *SolicitudesEntrenadorController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create SolicitudesEntrenador
// @Param	body		body 	models.SolicitudesEntrenador	true		"body for SolicitudesEntrenador content"
// @Success 201 {int} models.SolicitudesEntrenador
// @Failure 400 body is empty or invalid
// @router / [post]
func (c *SolicitudesEntrenadorController) Post() {
	var v models.SolicitudesEntrenador
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor: la solicitud contiene un parámetro incorrecto"}
	} else if _, err := models.AddSolicitudesEntrenador(&v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error al crear solicitud de entrenador", "error": err.Error()}
	} else {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = map[string]interface{}{"success": true, "status": 201, "message": "Solicitud creada correctamente", "data": v}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get SolicitudesEntrenador by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SolicitudesEntrenador
// @Failure 400 :id is invalid or not found
// @router /:id [get]
func (c *SolicitudesEntrenadorController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	v, err := models.GetSolicitudesEntrenadorById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "No se encontró la solicitud", "error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición correcta", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get SolicitudesEntrenador
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SolicitudesEntrenador
// @Failure 400
// @router / [get]
func (c *SolicitudesEntrenadorController) GetAll() {
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
				c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error: query inválida. Use formato clave:valor"}
				c.ServeJSON()
				return
			}
			query[kv[0]] = kv[1]
		}
	}

	l, err := models.GetAllSolicitudesEntrenador(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error al obtener solicitudes", "error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición correcta", "data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the SolicitudesEntrenador
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SolicitudesEntrenador	true		"body for SolicitudesEntrenador content"
// @Success 200 {object} models.SolicitudesEntrenador
// @Failure 400 :id is not int or body is invalid
// @router /:id [put]
func (c *SolicitudesEntrenadorController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	v := models.SolicitudesEntrenador{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor Put: la solicitud contiene un parámetro incorrecto"}
	} else if err := models.UpdateSolicitudesEntrenadorById(&v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor Put: no se pudo actualizar", "error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición correcta", "data": id}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the SolicitudesEntrenador
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 400 id is invalid or delete failed
// @router /:id [delete]
func (c *SolicitudesEntrenadorController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	if err := models.DeleteSolicitudesEntrenador(id); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor Delete: no se pudo eliminar", "error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Solicitud eliminada correctamente", "data": id}
	}
	c.ServeJSON()
}
