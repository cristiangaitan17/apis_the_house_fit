package models

import (
    "errors"
    "fmt"
    "reflect"
    "strings"
    "time"

    "github.com/beego/beego/v2/client/orm"
)

type DietaComidas struct {
    Id                int       `orm:"column(id);pk;auto"`
    DietaId           int       `orm:"column(dieta_id)"`
    TiempoComida      string    `orm:"column(tiempo_comida);null"`
    Descripcion       string    `orm:"column(descripcion);null"`
    Orden             int       `orm:"column(orden);null"`
    Activo            bool      `orm:"column(activo);default(true)"`
    FechaModificacion time.Time `orm:"column(Fecha_modificacion);auto_now;type(timestamp)"`
    FechaCreacion     time.Time `orm:"column(Fecha_creacion);auto_now_add;type(timestamp)"`
}

func (t *DietaComidas) TableName() string {
    return "dieta_comidas"
}

func init() {
    orm.RegisterModel(new(DietaComidas))
}

// AddDietaComidas insert a new DietaComidas into database
func AddDietaComidas(m *DietaComidas) (id int64, err error) {
    o := orm.NewOrm()
    id, err = o.Insert(m)
    return
}

// GetDietaComidasById retrieves DietaComidas by Id
func GetDietaComidasById(id int) (v *DietaComidas, err error) {
    o := orm.NewOrm()
    v = &DietaComidas{Id: id}
    if err = o.Read(v); err == nil {
        return v, nil
    }
    return nil, err
}

// GetAllDietaComidas retrieves all DietaComidas
func GetAllDietaComidas(query map[string]string, fields []string, sortby []string, order []string,
    offset int64, limit int64) (ml []interface{}, err error) {
    o := orm.NewOrm()
    qs := o.QueryTable(new(DietaComidas))
    
    for k, v := range query {
        k = strings.Replace(k, ".", "__", -1)
        if strings.Contains(k, "isnull") {
            qs = qs.Filter(k, (v == "true" || v == "1"))
        } else {
            qs = qs.Filter(k, v)
        }
    }
    
    var sortFields []string
    if len(sortby) != 0 {
        if len(sortby) == len(order) {
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

    var l []DietaComidas
    qs = qs.OrderBy(sortFields...)
    if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
        if len(fields) == 0 {
            for _, v := range l {
                ml = append(ml, v)
            }
        } else {
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

// UpdateDietaComidasById updates DietaComidas by Id
func UpdateDietaComidasById(m *DietaComidas) (err error) {
    o := orm.NewOrm()
    v := DietaComidas{Id: m.Id}
    if err = o.Read(&v); err == nil {
        var num int64
        if num, err = o.Update(m); err == nil {
            fmt.Println("Number of records updated in database:", num)
        }
    }
    return
}

// DeleteDietaComidas deletes DietaComidas by Id
func DeleteDietaComidas(id int) (err error) {
    o := orm.NewOrm()
    v := DietaComidas{Id: id}
    if err = o.Read(&v); err == nil {
        var num int64
        if num, err = o.Delete(&DietaComidas{Id: id}); err == nil {
            fmt.Println("Number of records deleted in database:", num)
        }
    }
    return
}