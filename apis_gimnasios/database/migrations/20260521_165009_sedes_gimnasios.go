package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type SedesGimnasios_20260521_165009 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &SedesGimnasios_20260521_165009{}
	m.Created = "20260521_165009"

	migration.Register("SedesGimnasios_20260521_165009", m)
}

// Run the migrations
func (m *SedesGimnasios_20260521_165009) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE sedes_gimnasios(id integer DEFAULT NULL,nombre TEXT NOT NULL,nit TEXT NOT NULL,descripcion TEXT NOT NULL,ciudad TEXT NOT NULL,departamento TEXT NOT NULL,direccion TEXT NOT NULL,correo TEXT NOT NULL,telefono TEXT NOT NULL,agregar_img TEXT NOT NULL,agregar__sede TEXT NOT NULL,aprovacion_entrenadores boolean NOT NULL,calificacion_prom numeric NOT NULL,total_resenas integer DEFAULT NULL,activo boolean NOT NULL,administrador_id integer DEFAULT NULL)")
}

// Reverse the migrations
func (m *SedesGimnasios_20260521_165009) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE sedes_gimnasios")
}
