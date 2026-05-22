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
	m.SQL("INSERT INTO login.login (id, username, password, fecha_creacion) VALUES (1, 'admin', 'admin123', NOW())")
	

}

// Reverse the migrations
func (m *InsertTablaLogin_20260521_172434) Down() {
	m.SQL("DELETE FROM login.login WHERE id = 1")

}
