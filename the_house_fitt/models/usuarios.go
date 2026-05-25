package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Usuarios struct {
	IdUsuario         int       `orm:"column(id_usuario);pk;auto"`
	Nombres           string    `orm:"column(nombres)"`
	Apellidos         string    `orm:"column(apellidos)"`
	Email             string    `orm:"column(email)"`
	FechaNacimiento   time.Time `orm:"column(Fecha_nacimiento);type(timestamp without time zone);null"`
	Nacionalidad      string    `orm:"column(nacionalidad)"`
	Ciudad            string    `orm:"column(ciudad)"`
	Departamento      string    `orm:"column(departamento)"`
	Direccion         string    `orm:"column(direccion)"`
	Telefono          string    `orm:"column(telefono)"`
	Activo            bool      `orm:"column(activo)"`
	FechaModificacion time.Time `orm:"column(Fecha_modificacion);type(timestamp without time zone);null;auto_now"`
	FechaCreacion     time.Time `orm:"column(Fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
}

func (t *Usuarios) TableName() string {
	return "LOGIN.usuarios"
}

func init() {
	orm.RegisterModel(new(Usuarios))
}