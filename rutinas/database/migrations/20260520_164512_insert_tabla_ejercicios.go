package migrations

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaEjercicios_20260520_164512 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaEjercicios_20260520_164512{}
	m.Created = "20260520_164512"

	migration.Register("InsertTablaEjercicios_20260520_164512", m)
}

// Run the migrations
func (m *InsertTablaEjercicios_20260520_164512) Up() {
	m.SQL("INSERT INTO rutinas.ejercicios (grupo_muscular_id, nombre, descripcion_corta, descripcion_larga, posicion_inicial, ejecucion, consejos, nivel) VALUES (1, 'Press de banca', 'Ejercicio de fuerza para el pecho', 'Ejercicio que consiste en empujar un peso hacia arriba desde el pecho', 'acostado', 'realizar press de banca', 'mantener la espalda plana', 1)")

}

// Reverse the migrations
func (m *InsertTablaEjercicios_20260520_164512) Down() {
	m.SQL("DELETE FROM rutinas.ejercicios WHERE nombre='Press de banca'") 
}
