package migrations

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaGruposMusculares_20260521_144559 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaGruposMusculares_20260521_144559{}
	m.Created = "20260521_144559"

	migration.Register("InsertTablaGruposMusculares_20260521_144559", m)
}

// Run the migrations
func (m *InsertTablaGruposMusculares_20260521_144559) Up() {
	m.SQL("INSERT INTO rutinas.grupos_musculares (nombre, descripcion, activo) VALUES ('Pecho', 'Grupo muscular del torso que se encuentra en la parte frontal del cuerpo', true)")
}

// Reverse the migrations
func (m *InsertTablaGruposMusculares_20260521_144559) Down() {
	m.SQL("DELETE FROM rutinas.grupos_musculares WHERE nombre='Pecho'")

}
