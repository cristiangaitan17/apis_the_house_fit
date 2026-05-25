package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type SolicitudesEntrenador struct {
	Id                int       `orm:"column(id);pk;auto" json:"id"`
	UsuarioId         int       `orm:"column(usuario_id)" json:"usuario_id"`
	GimnasioId        int       `orm:"column(gimnasio_id)" json:"gimnasio_id"`
	NombresApellidos  string    `orm:"column(nombres_apellidos);null" json:"nombres_apellidos,omitempty"`
	Cedula            string    `orm:"column(cedula);null" json:"cedula,omitempty"`
	Experiencia       string    `orm:"column(experiencia);null" json:"experiencia,omitempty"`
	SobreMi           string    `orm:"column(sobre_mi);null" json:"sobre_mi,omitempty"`
	Especialidad      string    `orm:"column(especialidad);null" json:"especialidad,omitempty"`
	Whatsapp          string    `orm:"column(whatsapp);null" json:"whatsapp,omitempty"`
	Correo            string    `orm:"column(correo);null" json:"correo,omitempty"`
	Direccion         string    `orm:"column(direccion);null" json:"direccion,omitempty"`
	Estado            string    `orm:"column(estado);null" json:"estado,omitempty"`
	RevisadoPor       bool      `orm:"column(revisado_por);null" json:"revisado_por,omitempty"`
	RevisadoEn        time.Time `orm:"column(revisado_en);type(timestamp without time zone);null;auto_now" json:"revisado_en,omitempty"`
	CreadoEn          time.Time `orm:"column(creado_en);type(timestamp without time zone);null;auto_now_add" json:"creado_en,omitempty"`
	Activo            bool      `orm:"column(activo);default(true)" json:"activo"`
	FechaModificacion time.Time `orm:"column(Fecha_modificacion);type(timestamp without time zone);null;auto_now" json:"fecha_modificacion,omitempty"`
	FechaCreacion     time.Time `orm:"column(Fecha_creacion);type(timestamp without time zone);null;auto_now_add" json:"fecha_creacion,omitempty"`
}

func (t *SolicitudesEntrenador) TableName() string {
	return "solicitudes_entrenador"
}

func init() {
	orm.RegisterModel(new(SolicitudesEntrenador))
}

// AddSolicitudesEntrenador insert a new SolicitudesEntrenador into database and returns
// last inserted Id on success.
func AddSolicitudesEntrenador(m *SolicitudesEntrenador) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSolicitudesEntrenadorById retrieves SolicitudesEntrenador by Id. Returns error if
// Id doesn't exist
func GetSolicitudesEntrenadorById(id int) (v *SolicitudesEntrenador, err error) {
	o := orm.NewOrm()
	v = &SolicitudesEntrenador{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudesEntrenador retrieves all SolicitudesEntrenador matches certain condition. Returns empty list if
// no records exist
func GetAllSolicitudesEntrenador(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudesEntrenador))
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

	var l []SolicitudesEntrenador
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

// UpdateSolicitudesEntrenador updates SolicitudesEntrenador by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudesEntrenadorById(m *SolicitudesEntrenador) (err error) {
	o := orm.NewOrm()
	v := SolicitudesEntrenador{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSolicitudesEntrenador deletes SolicitudesEntrenador by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudesEntrenador(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudesEntrenador{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SolicitudesEntrenador{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
