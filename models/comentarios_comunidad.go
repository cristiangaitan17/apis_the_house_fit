package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ComentariosComunidad struct {
	Id                int       `orm:"column(id);pk"`
	CategoriaId       int       `orm:"column(categoria_id)"`
	UsuarioId         int       `orm:"column(usuario_id)"`
	Contenido         string    `orm:"column(contenido)"`
	Calificacion      int       `orm:"column(calificacion);null"`
	Likes             int       `orm:"column(likes);null"`
	Dislikes          int       `orm:"column(dislikes);null"`
	Estado            string    `orm:"column(estado);null"`
	Activo            bool      `orm:"column(activo)"`
	FechaModificacion time.Time `orm:"column(Fecha_modificacion);type(timestamp without time zone)"`
	FechaCreacion     time.Time `orm:"column(Fecha_creacion);type(timestamp without time zone)"`
}

func (t *ComentariosComunidad) TableName() string {
	return "comentarios_comunidad"
}

func init() {
	orm.RegisterModel(new(ComentariosComunidad))
}

// AddComentariosComunidad insert a new ComentariosComunidad into database and returns
// last inserted Id on success.
func AddComentariosComunidad(m *ComentariosComunidad) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetComentariosComunidadById retrieves ComentariosComunidad by Id. Returns error if
// Id doesn't exist
func GetComentariosComunidadById(id int) (v *ComentariosComunidad, err error) {
	o := orm.NewOrm()
	v = &ComentariosComunidad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllComentariosComunidad retrieves all ComentariosComunidad matches certain condition. Returns empty list if
// no records exist
func GetAllComentariosComunidad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ComentariosComunidad))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ComentariosComunidad
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateComentariosComunidad updates ComentariosComunidad by Id and returns error if
// the record to be updated doesn't exist
func UpdateComentariosComunidadById(m *ComentariosComunidad) (err error) {
	o := orm.NewOrm()
	v := ComentariosComunidad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteComentariosComunidad deletes ComentariosComunidad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteComentariosComunidad(id int) (err error) {
	o := orm.NewOrm()
	v := ComentariosComunidad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ComentariosComunidad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
