package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Categorias_20260521_134524 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Categorias_20260521_134524{}
	m.Created = "20260521_134524"

	migration.Register("Categorias_20260521_134524", m)
}

// Run the migrations
func (m *Categorias_20260521_134524) Up() {
	file, err:=ioutil.ReadFile("../scripts/20260521_134524_categorias_up.sql")

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
func (m *Categorias_20260521_134524) Down() {
	 file, err:=ioutil.ReadFile("../scripts/20260521_134524_categorias_down.sql")

	if err != nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file),";")
	for _, request:=range requests{
		fmt.Println(requests)
		m.SQL(request)
	}
}
