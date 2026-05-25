package migrations

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaRutinas_20260521_144624 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaRutinas_20260521_144624{}
	m.Created = "20260521_144624"

	migration.Register("InsertTablaRutinas_20260521_144624", m)
}

// Run the migrations
func (m *InsertTablaRutinas_20260521_144624) Up() {
	m.SQL("INSERT INTO rutinas.rutinas (nombre, descripcion, objetivo, activo) VALUES ('Rutina de fuerza', 'Rutina enfocada en el desarrollo de la fuerza muscular', 'Desarrollar la fuerza muscular a través de ejercicios compuestos y progresión de cargas', true)")

}

// Reverse the migrations
func (m *InsertTablaRutinas_20260521_144624) Down() {
	m.SQL("DELETE FROM rutinas.rutinas WHERE nombre='Rutina de fuerza'")
}
