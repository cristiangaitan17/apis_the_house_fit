package controllers

import (
	"apis_gimnasios/models"
	"encoding/json"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
)

// GimnasioClasesController operations for GimnasioClases
type GimnasioClasesController struct {
	beego.Controller
}

// URLMapping ...
func (c *GimnasioClasesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
func (c *GimnasioClasesController) Post() {
	var v models.GimnasioClases
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddGimnasioClases(&v); err == nil {
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
func (c *GimnasioClasesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetGimnasioClasesById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 404, "message": "Registro no encontrado"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición exitosa", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
func (c *GimnasioClasesController) GetAll() {
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
				c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error: invalid query key/value pair"}
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllGimnasioClases(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor GetAll: la solicitud contiene un parámetro incorrecto o no contiene registro solicitado"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición exitosa", "data": l}
	}
	c.ServeJSON()
}

// Put ...
func (c *GimnasioClasesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.GimnasioClases{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateGimnasioClasesById(&v); err == nil {
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
func (c *GimnasioClasesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteGimnasioClases(id); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Registro eliminado exitosamente"}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 404, "message": "Registro no encontrado"}
	}
	c.ServeJSON()
}