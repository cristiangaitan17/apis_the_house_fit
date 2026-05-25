package controllers

import (
	"entrenadores/models"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// EntrenadorDocumentosController operations for EntrenadorDocumentos
type EntrenadorDocumentosController struct {
	beego.Controller
}

// URLMapping ...
func (c *EntrenadorDocumentosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create EntrenadorDocumentos
// @Param	body		body 	models.EntrenadorDocumentos	true		"body for EntrenadorDocumentos content"
// @Success 201 {int} models.EntrenadorDocumentos
// @Failure 400 body is empty or invalid
// @router / [post]
func (c *EntrenadorDocumentosController) Post() {
	var v models.EntrenadorDocumentos
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor: la solicitud contiene un parámetro incorrecto"}
	} else if _, err := models.AddEntrenadorDocumentos(&v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error al crear documento del entrenador", "error": err.Error()}
	} else {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = map[string]interface{}{"success": true, "status": 201, "message": "Documento creado correctamente", "data": v}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get EntrenadorDocumentos by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EntrenadorDocumentos
// @Failure 400 :id is invalid or not found
// @router /:id [get]
func (c *EntrenadorDocumentosController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	v, err := models.GetEntrenadorDocumentosById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "No se encontró el documento", "error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición correcta", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get EntrenadorDocumentos
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.EntrenadorDocumentos
// @Failure 400
// @router / [get]
func (c *EntrenadorDocumentosController) GetAll() {
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

	l, err := models.GetAllEntrenadorDocumentos(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error al obtener documentos", "error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición correcta", "data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the EntrenadorDocumentos
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EntrenadorDocumentos	true		"body for EntrenadorDocumentos content"
// @Success 200 {object} models.EntrenadorDocumentos
// @Failure 400 :id is not int or body is invalid
// @router /:id [put]
func (c *EntrenadorDocumentosController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	v := models.EntrenadorDocumentos{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor Put: la solicitud contiene un parámetro incorrecto"}
	} else if err := models.UpdateEntrenadorDocumentosById(&v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor Put: no se pudo actualizar", "error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición correcta", "data": id}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the EntrenadorDocumentos
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 400 id is invalid or delete failed
// @router /:id [delete]
func (c *EntrenadorDocumentosController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	if err := models.DeleteEntrenadorDocumentos(id); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor Delete: no se pudo eliminar", "error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Documento eliminado correctamente", "data": id}
	}
	c.ServeJSON()
}
