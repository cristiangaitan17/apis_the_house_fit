package migrations

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaUsuarios_20260521_172444 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaUsuarios_20260521_172444{}
	m.Created = "20260521_172444"

	migration.Register("InsertTablaUsuarios_20260521_172444", m)
}

// Run the migrations
func (m *InsertTablaUsuarios_20260521_172444) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTablaUsuarios_20260521_172444) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
