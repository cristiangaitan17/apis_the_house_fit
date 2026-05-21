package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Noticias_20260521_134711 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Noticias_20260521_134711{}
	m.Created = "20260521_134711"

	migration.Register("Noticias_20260521_134711", m)
}

// Run the migrations
func (m *Noticias_20260521_134711) Up() {
	file, err:=ioutil.ReadFile("../scripts/20260521_134711_noticias_up.sql")

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
func (m *Noticias_20260521_134711) Down() {
	 file, err:=ioutil.ReadFile("../scripts/20260521_134711_noticias_down.sql")

	if err != nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file),";")
	for _, request:=range requests{
		fmt.Println(requests)
		m.SQL(request)
	}
}

