package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Noticias struct {
	Id                int       `orm:"column(id);pk;auto"`
	CategoriaId       *Categorias `orm:"column(categoria_id);rel(fk)"`
	Titulo            string    `orm:"column(titulo)"`
	Contenido         string    `orm:"column(contenido)"`
	Encabezado        string    `orm:"column(encabezado);null"`
	ImagenPrincipal   string    `orm:"column(imagen_principal);null"`
	AutorId           int       `orm:"column(autor_id)"`
	Estado            string    `orm:"column(estado);null"`
	Vistas            int       `orm:"column(vistas);null"`
	PublicadoEn       time.Time `orm:"column(publicado_en);type(timestamp without time zone);null"`
	CreadoEn          time.Time `orm:"column(creado_en);type(timestamp without time zone);null"`
	ActualizadoEn     time.Time `orm:"column(actualizado_en);type(timestamp without time zone);null"`
	Activo            bool      `orm:"column(activo)"`
	fechac_creacion     time.Time  `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
	fecha_actualizacion time.Time  `orm:"column(fecha_modificacion);type(timestamp without time zone);null;auto_now"`
}

func (n *Noticias) TableName() string {
    return "noticias"  // ← minúsculas, no "Noticias"
}

func init() {
	orm.RegisterModel(new(Noticias))
}

// AddNoticias insert a new Noticias into database and returns
// last inserted Id on success.
func AddNoticias(m *Noticias) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNoticiasById obtiene una noticia con su categoría relacionada
func GetNoticiasById(id int) (v *Noticias, err error) {
	o := orm.NewOrm()
	v = &Noticias{Id: id}
	if err = o.Read(v); err == nil {
		// Cargar la categoría relacionada
		o.LoadRelated(v, "CategoriaId")
		return v, nil
	}
	return nil, err
}

// GetAllNoticias retrieves all Noticias matches certain condition. Returns empty list if
// no records exist
func GetAllNoticias(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Noticias)).RelatedSel()
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

	var l []Noticias
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

// UpdateNoticias updates Noticias by Id and returns error if
// the record to be updated doesn't exist
func UpdateNoticiasById(m *Noticias) (err error) {
	o := orm.NewOrm()
	v := Noticias{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNoticias deletes Noticias by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNoticias(id int) (err error) {
	o := orm.NewOrm()
	v := Noticias{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Noticias{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
