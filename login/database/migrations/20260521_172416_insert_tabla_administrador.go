package migrations

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaAdministrador_20260521_172416 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaAdministrador_20260521_172416{}
	m.Created = "20260521_172416"

	migration.Register("InsertTablaAdministrador_20260521_172416", m)
}

// Run the migrations
func (m *InsertTablaAdministrador_20260521_172416) Up() {
	m.SQL("INSERT INTO rutinas.ejercicios (nombre_gym, nit, anio_fundacion, cantidad_clientes, departamento, ciudad, direccion, propietario_nombre, telefono, correo, pagina_web, firma, fecha_firma) VALUES ('Gimnasio XYZ', '123456789', 2020, 100, 'Bogotá', 'Bogotá', 'Calle 123', 'Juan Pérez', '1234567890', 'juan@ejemplo.com', 'https://www.gimnasioxyz.com', 'firma_juan.pdf', '2023-01-01')")

}
// Reverse the migrations
func (m *InsertTablaAdministrador_20260521_172416) Down() {
	m.SQL("DELETE FROM rutinas.ejercicios WHERE nit = '123456789'")

}
