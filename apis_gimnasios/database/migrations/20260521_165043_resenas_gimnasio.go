package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ResenasGimnasio_20260521_165043 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ResenasGimnasio_20260521_165043{}
	m.Created = "20260521_165043"

	migration.Register("ResenasGimnasio_20260521_165043", m)
}

// Run the migrations
func (m *ResenasGimnasio_20260521_165043) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE resenas_gimnasio(id integer DEFAULT NULL,gimnasio_id integer DEFAULT NULL,usuario_id integer DEFAULT NULL,calificacion integer DEFAULT NULL,comentario TEXT NOT NULL,activo boolean NOT NULL)")
}

// Reverse the migrations
func (m *ResenasGimnasio_20260521_165043) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE resenas_gimnasio")
}
