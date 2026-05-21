package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Ejercicios struct {
	Id                int       `orm:"column(id);pk"`
	GrupoMuscularId   int       `orm:"column(grupo_muscular_id)"`
	Nombre            string    `orm:"column(nombre)"`
	DescripcionCorta  string    `orm:"column(descripcion_corta);null"`
	DescripcionLarga  string    `orm:"column(descripcion_larga);null"`
	PosicionInicial   string    `orm:"column(posicion_inicial);null"`
	Ejecucion         string    `orm:"column(ejecucion);null"`
	Consejos          string    `orm:"column(consejos);null"`
	Nivel             string    `orm:"column(nivel);null"`
	Activo            bool      `orm:"column(activo)"`
	FechaModificacion time.Time `orm:"column(Fecha_modificacion);type(timestamp without time zone)"`
	FechaCreacion     time.Time `orm:"column(Fecha_creacion);type(timestamp without time zone)"`
}

func (t *Ejercicios) TableName() string {
	return "ejercicios"
}

func init() {
	orm.RegisterModel(new(Ejercicios))
}

// AddEjercicios insert a new Ejercicios into database and returns
// last inserted Id on success.
func AddEjercicios(m *Ejercicios) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEjerciciosById retrieves Ejercicios by Id. Returns error if
// Id doesn't exist
func GetEjerciciosById(id int) (v *Ejercicios, err error) {
	o := orm.NewOrm()
	v = &Ejercicios{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEjercicios retrieves all Ejercicios matches certain condition. Returns empty list if
// no records exist
func GetAllEjercicios(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Ejercicios))
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

	var l []Ejercicios
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

// UpdateEjercicios updates Ejercicios by Id and returns error if
// the record to be updated doesn't exist
func UpdateEjerciciosById(m *Ejercicios) (err error) {
	o := orm.NewOrm()
	v := Ejercicios{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEjercicios deletes Ejercicios by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEjercicios(id int) (err error) {
	o := orm.NewOrm()
	v := Ejercicios{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Ejercicios{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
