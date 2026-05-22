package migrations

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaLogin_20260521_172434 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaLogin_20260521_172434{}
	m.Created = "20260521_172434"

	migration.Register("InsertTablaLogin_20260521_172434", m)
}

// Run the migrations
func (m *InsertTablaLogin_20260521_172434) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTablaLogin_20260521_172434) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
