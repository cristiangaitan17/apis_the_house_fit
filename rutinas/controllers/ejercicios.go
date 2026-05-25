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

// EjerciciosController operations for Ejercicios
type EjerciciosController struct {
	beego.Controller
}

// URLMapping ...
func (c *EjerciciosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Ejercicios
// @Param	body		body 	models.Ejercicios	true		"body for Ejercicios content"
// @Success 201 {int} models.Ejercicios
// @Failure 403 body is empty
// @router / [post]
func (c *EjerciciosController) Post() {
	var v models.Ejercicios
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor: la solicitud contiene un parametro incorrecto"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "peticion correcta", "data": v}
	}
	c.ServeJSON()

}

// GetOne ...
// @Title Get One
// @Description get Ejercicios by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Ejercicios
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EjerciciosController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEjerciciosById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor GetOne: la solicitud contiene un parametro incorrecto o no contiene registro solicitado"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "peticion correcta", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Ejercicios
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Ejercicios
// @Failure 403
// @router / [get]
func (c *EjerciciosController) GetAll() {
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

	l, err := models.GetAllEjercicios(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor GetAll: la solicitud contiene un parametro incorrecto o no contiene registro solicitado"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "peticion correcta", "data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Ejercicios
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Ejercicios	true		"body for Ejercicios content"
// @Success 200 {object} models.Ejercicios
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EjerciciosController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Ejercicios{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor Put: la solicitud contiene un parametro incorrecto"}
	} else {
		if err := models.UpdateEjerciciosById(&v); err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "Error en el servidor Put: no se pudo actualizar"}
		} else {
			c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "message": "peticion correcta", "data": id}
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Ejercicios
// @Param	id		path 	int	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 400 Id inválido o registro no encontrado
// @router /:id [delete]
func (c *EjerciciosController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	if idStr == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "No se recibió ningún ID para eliminar",
		}
		c.ServeJSON()
		return
	}

	// 2. Convertir el parámetro a un número entero
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El ID proporcionado no es un formato numérico válido",
		}
		c.ServeJSON()
		return
	}

	// 3. Llamar al modelo para ejecutar el borrado físico o lógico en Postgres
	if err := models.DeleteEjercicios(id); err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "Error en el servidor Delete: la solicitud contiene un parametro incorrecto o el registro no existe",
		}
		c.ServeJSON()
		return
	}

	// 4. Respuesta exitosa
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"status":  200,
		"message": "El ejercicio ha sido eliminado correctamente de la base de datos",
	}
	c.ServeJSON()
}
