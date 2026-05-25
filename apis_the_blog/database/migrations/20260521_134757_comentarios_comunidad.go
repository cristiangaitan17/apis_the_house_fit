package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ComentariosComunidad_20260521_134757 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ComentariosComunidad_20260521_134757{}
	m.Created = "20260521_134757"

	migration.Register("ComentariosComunidad_20260521_134757", m)
}

// Run the migrations
func (m *ComentariosComunidad_20260521_134757) Up() {
	file, err:=ioutil.ReadFile("../scripts/20260521_134757_comentarios_comunidad_up.sql")

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
func (m *ComentariosComunidad_20260521_134757) Down() {
	 file, err:=ioutil.ReadFile("../scripts/20260521_134757_comentarios_comunidad_down.sql")

	if err != nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file),";")
	for _, request:=range requests{
		fmt.Println(requests)
		m.SQL(request)
	}
}

