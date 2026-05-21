package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Pedidos struct {
	Id                int       `orm:"column(id_pedido);pk;auto"`
	UsuarioId         int       `orm:"column(usuario_id)"`
	Subtotal          float64   `orm:"column(subtotal)"`
	Envio             float64   `orm:"column(envio)"`
	Impuesto          float64   `orm:"column(impuesto)"`
	Total             float64   `orm:"column(total)"`
	MetodoPago        string    `orm:"column(metodo_pago)"`
	EstadoPago        string    `orm:"column(estado_pago)"`
	EstadoPedido      string    `orm:"column(estado_pedido)"`
	NombreReceptor    string    `orm:"column(nombre_receptor)"`
	ApellidoReceptor  string    `orm:"column(apellido_receptor)"`
	EmailReceptor     string    `orm:"column(email_receptor)"`
	Telefono          string    `orm:"column(telefono)"`
	DireccionEnvio    string    `orm:"column(direccion_envio)"`
	Activo            bool      `orm:"column(activo)"`
	FechaModificacion time.Time `orm:"column(Fecha_modificacion);type(timestamp without time zone);null;auto_now"`
	FechaCreacion     time.Time `orm:"column(Fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
}


func (t *Pedidos) TableName() string {
	return "pedidos"
}

func init() {
	orm.RegisterModel(new(Pedidos))
}

// AddPedidos insert a new Pedidos into database and returns
// last inserted Id on success.
func AddPedidos(m *Pedidos) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPedidosById retrieves Pedidos by Id. Returns error if
// Id doesn't exist
func GetPedidosById(id int) (v *Pedidos, err error) {
	o := orm.NewOrm()
	v = &Pedidos{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPedidos retrieves all Pedidos matches certain condition. Returns empty list if
// no records exist
func GetAllPedidos(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Pedidos))
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

	var l []Pedidos
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

// UpdatePedidos updates Pedidos by Id and returns error if
// the record to be updated doesn't exist
func UpdatePedidosById(m *Pedidos) (err error) {
	o := orm.NewOrm()
	v := Pedidos{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePedidos deletes Pedidos by Id and returns error if
// the record to be deleted doesn't exist
func DeletePedidos(id int) (err error) {
	o := orm.NewOrm()
	v := Pedidos{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Pedidos{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
