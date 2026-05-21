package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type DietaComidas_20260521_134747 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &DietaComidas_20260521_134747{}
	m.Created = "20260521_134747"

	migration.Register("DietaComidas_20260521_134747", m)
}

// Run the migrations
func (m *DietaComidas_20260521_134747) Up() {
	file, err:=ioutil.ReadFile("../scripts/20260521_134747_dieta_comidas_up.sql")

	if err != nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file),";")
	for _, request:=range requests{
		fmt.Println(requests)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *DietaComidas_20260521_134747) Down() {
	 file, err:=ioutil.ReadFile("../scripts/20260521_134747_dieta_comidas_down.sql")

	if err != nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file),";")
	for _, request:=range requests{
		fmt.Println(requests)
		m.SQL(request)
	}
}
