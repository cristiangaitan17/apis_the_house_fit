package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaEntrenadorDocumentos_20260524_212143 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaEntrenadorDocumentos_20260524_212143{}
	m.Created = "20260524_212143"

	migration.Register("InsertTablaEntrenadorDocumentos_20260524_212143", m)
}

// Run the migrations
func (m *InsertTablaEntrenadorDocumentos_20260524_212143) Up() {
	m.SQL("INSERT INTO entrenadores.documentos (entrenador_id, identificacion, nombre_archivo, creado_en) VALUES (1, '123456789', 'documento.pdf', NOW())")

}

// Reverse the migrations
func (m *InsertTablaEntrenadorDocumentos_20260524_212143) Down() {
	m.SQL("DELETE FROM entrenadores.documentos WHERE entrenador_id = 1 AND identificacion = '123456789' AND nombre_archivo = 'documento.pdf'")
}
