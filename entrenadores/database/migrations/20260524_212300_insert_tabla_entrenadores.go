package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaEntrenadores_20260524_212300 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaEntrenadores_20260524_212300{}
	m.Created = "20260524_212300"

	migration.Register("InsertTablaEntrenadores_20260524_212300", m)
}

// Run the migrations
func (m *InsertTablaEntrenadores_20260524_212300) Up() {
	m.SQL("INSERT INTO entrenadores.entrenadores (id, id_administrador, cedula, descripcion, especialidad, gimnasio_id, anios_experiencia, foto_url, aprovacion_entrenador, hoja_vida, calificacion_prom) VALUES (1, 1, '123456789', 'Descripcion del entrenador', 'Especialidad del entrenador', 1, 5, 'http://example.com/foto.jpg', 'Aprovacion_entrenador', 'Hoja_vida', 4.5)")
}

// Reverse the migrations
func (m *InsertTablaEntrenadores_20260524_212300) Down() {
	m.SQL("DELETE FROM entrenadores.entrenadores WHERE id = 1")
}
