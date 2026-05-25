package main
import (
	"fmt"
	"strings"
	"io/ioutil"
	
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY

type CreacionDb_20260521_125236 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDb_20260521_125236{}
	m.Created = "20260521_125236"
							
	migration.Register("CreacionDb_20260521_125236", m)
}
					   
// Run the migrations
func (m *CreacionDb_20260521_125236) Up() {
	file, err := ioutil.ReadFile("database/scripts/20260514_creacion_db_up.sql")

	if err != nil {
		fmt.Println(err)
		
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		fmt.Println(request)
			m.SQL(request)
	}
}




		// use m.SQL("CREATE TABLE ...") to make schema update
					

// Reverse the migrations
func (m *CreacionDb_20260521_125236) Down() {

	file, err := ioutil.ReadFile("database/scripts/20260514_creacion_db_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		fmt.Println(request)
			m.SQL(request)
	}
// use m.SQL("DROP TABLE ...") to reverse schema update
					
}
				