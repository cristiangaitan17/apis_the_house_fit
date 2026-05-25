package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type GimnasioClases_20260521_165035 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &GimnasioClases_20260521_165035{}
	m.Created = "20260521_165035"

	migration.Register("GimnasioClases_20260521_165035", m)
}

// Run the migrations
func (m *GimnasioClases_20260521_165035) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE gimnasio_clases(id integer DEFAULT NULL,gimnasio_id integer DEFAULT NULL,nombre_clase TEXT NOT NULL,activo boolean NOT NULL)")
}

// Reverse the migrations
func (m *GimnasioClases_20260521_165035) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE gimnasio_clases")
}
