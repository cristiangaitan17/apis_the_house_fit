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

// GruposMuscularesController operations for GruposMusculares
type GruposMuscularesController struct {
	beego.Controller
}

// URLMapping ...
func (c *GruposMuscularesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create GruposMusculares
// @Param	body		body 	models.GruposMusculares	true		"body for GruposMusculares content"
// @Success 201 {int} models.GruposMusculares
// @Failure 403 body is empty
// @router / [post]
func (c *GruposMuscularesController) Post() {
	var v models.GruposMusculares
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor: la solicitud contiene un parametro incorrecto"}
	} else {
		if _, err := models.AddGruposMusculares(&v); err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error en el servidor: no se pudo crear el grupo muscular"}
		} else {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"success": true, "status": 201, "message": "Grupo muscular creado", "data": v}
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get GruposMusculares by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GruposMusculares
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GruposMuscularesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "El ID proporcionado no es válido"}
		c.ServeJSON()
		return
	}
	v, err := models.GetGruposMuscularesById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 404, "message": "Grupo muscular no encontrado"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición correcta", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get GruposMusculares
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.GruposMusculares
// @Failure 403
// @router / [get]
func (c *GruposMuscularesController) GetAll() {
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

	l, err := models.GetAllGruposMusculares(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor GetAll: parámetros inválidos o consulta vacía"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición correcta", "data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the GruposMusculares
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GruposMusculares	true		"body for GruposMusculares content"
// @Success 200 {object} models.GruposMusculares
// @Failure 403 :id is not int
// @router /:id [put]
func (c *GruposMuscularesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "El ID proporcionado no es válido"}
		c.ServeJSON()
		return
	}
	v := models.GruposMusculares{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor Put: la solicitud contiene un parametro incorrecto"}
	} else {
		if err := models.UpdateGruposMuscularesById(&v); err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor Put: no se pudo actualizar"}
		} else {
			c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Grupo muscular actualizado", "data": v}
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the GruposMusculares
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *GruposMuscularesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	if idStr == "" {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "No se recibió ningún ID para eliminar"}
		c.ServeJSON()
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "El ID proporcionado no es un formato numérico válido"}
		c.ServeJSON()
		return
	}
	if err := models.DeleteGruposMusculares(id); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor Delete: la solicitud contiene un parametro incorrecto o el registro no existe"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "El grupo muscular ha sido eliminado correctamente"}
	}
	c.ServeJSON()
}
