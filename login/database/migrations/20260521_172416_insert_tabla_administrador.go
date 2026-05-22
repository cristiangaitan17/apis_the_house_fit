package migrations

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaAdministrador_20260521_172416 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaAdministrador_20260521_172416{}
	m.Created = "20260521_172416"

	migration.Register("InsertTablaAdministrador_20260521_172416", m)
}

// Run the migrations
func (m *InsertTablaAdministrador_20260521_172416) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTablaAdministrador_20260521_172416) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
