package migrations

import (
	"strings"
	"fmt"
	"io/ioutil"
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionDb_20260520_163845 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDb_20260520_163845{}
	m.Created = "20260520_163845"

	migration.Register("CreacionDb_20260520_163845", m)
}

// Run the migrations
func (m *CreacionDb_20260520_163845) Up() {
		file, err:=ioutil.ReadFile("../scripts/20260520_163845_creacion_db_up.sql")
	if err != nil {
		fmt.Println(err)
	}
	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}

}

// Reverse the migrations
func (m *CreacionDb_20260520_163845) Down() {
			file, err:=ioutil.ReadFile("../scripts/20260520_163845_creacion_db_down.sql")
	if err != nil {
		fmt.Println(err)
	}
	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}


}
