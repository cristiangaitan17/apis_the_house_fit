package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaSolicitudesEntrenadores_20260524_212426 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaSolicitudesEntrenadores_20260524_212426{}
	m.Created = "20260524_212426"

	migration.Register("InsertTablaSolicitudesEntrenadores_20260524_212426", m)
}

// Run the migrations
func (m *InsertTablaSolicitudesEntrenadores_20260524_212426) Up() {
	m.SQL("INSERT INTO entrenadores.solicitudes_entrenadores (id, usuario_id, gimnasio_id, nombres_apellidos, cedula, telefono, correo_electronico, estado_solicitud, revisado_en, creado_en, activo, fecha_modificacion, fecha_creacion) VALUES (1, 1, 1, 'Juan Perez', '123456789', '555-1234', 'juan.perez@example.com', 'Pendiente', NULL, NOW(), TRUE, NULL, NOW())")

}

// Reverse the migrations
func (m *InsertTablaSolicitudesEntrenadores_20260524_212426) Down() {
	m.SQL("DELETE FROM entrenadores.solicitudes_entrenadores WHERE id = 1")
}
