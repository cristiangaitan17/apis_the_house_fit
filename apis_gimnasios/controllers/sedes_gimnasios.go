package controllers

import (
	"apis_gimnasios/models"
	"encoding/json"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
)

type SedesGimnasiosController struct {
	beego.Controller
}

func (c *SedesGimnasiosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
func (c *SedesGimnasiosController) Post() {
	var v models.SedesGimnasios
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSedesGimnasios(&v); err == nil {
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
func (c *SedesGimnasiosController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	
	v, err := models.GetSedesGimnasiosById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 404, "message": "Registro no encontrado"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición exitosa", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
func (c *SedesGimnasiosController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
	var offset int64 = 0

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
			k, val := kv[0], kv[1]
			query[k] = val
		}
	}

	l, err := models.GetAllSedesGimnasios(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Petición exitosa", "data": l}
	}
	c.ServeJSON()
}

// Put ...
func (c *SedesGimnasiosController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	
	var v models.SedesGimnasios
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": err.Error()}
		c.ServeJSON()
		return
	}
	
	v.Id = id
	if err := models.UpdateSedesGimnasiosById(&v); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Registro actualizado exitosamente", "data": v}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": err.Error()}
	}
	c.ServeJSON()
}

// Delete ...
func (c *SedesGimnasiosController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	
	if err := models.DeleteSedesGimnasios(id); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "Registro eliminado exitosamente"}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 404, "message": "Registro no encontrado"}
	}
	c.ServeJSON()
}