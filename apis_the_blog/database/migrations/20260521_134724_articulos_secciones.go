package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ArticulosSecciones_20260521_134724 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ArticulosSecciones_20260521_134724{}
	m.Created = "20260521_134724"

	migration.Register("ArticulosSecciones_20260521_134724", m)
}

// Run the migrations
func (m *ArticulosSecciones_20260521_134724) Up() {
	file, err:=ioutil.ReadFile("../scripts/20260521_134724_articulos_secciones_up.sql")

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
func (m *ArticulosSecciones_20260521_134724) Down() {
	 file, err:=ioutil.ReadFile("../scripts/20260521_134724_articulos_secciones_down.sql")

	if err != nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file),";")
	for _, request:=range requests{
		fmt.Println(requests)
		m.SQL(request)
	}
}
