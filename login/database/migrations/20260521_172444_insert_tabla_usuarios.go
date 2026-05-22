package migrations

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaUsuarios_20260521_172444 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaUsuarios_20260521_172444{}
	m.Created = "20260521_172444"

	migration.Register("InsertTablaUsuarios_20260521_172444", m)
}

// Run the migrations
func (m *InsertTablaUsuarios_20260521_172444) Up() {
	m.SQL("INSERT INTO login.usuarios (nombres, apellidos, email, fecha_nacimiento, nacionalidad, ciudad, departamento, direccion, telefono) VALUES ('Juan', 'Pérez', 'juan@ejemplo.com', '1990-01-01', 'Colombiana', 'Bogotá', 'Bogotá', 'Calle 123', '1234567890')")

}

// Reverse the migrations
func (m *InsertTablaUsuarios_20260521_172444) Down() {
	m.SQL("DELETE FROM login.usuarios WHERE email = 'juan@ejemplo.com'")

}
