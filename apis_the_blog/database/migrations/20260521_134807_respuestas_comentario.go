package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type RespuestasComentario_20260521_134807 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RespuestasComentario_20260521_134807{}
	m.Created = "20260521_134807"

	migration.Register("RespuestasComentario_20260521_134807", m)
}

// Run the migrations
func (m *RespuestasComentario_20260521_134807) Up() {
	file, err:=ioutil.ReadFile("../scripts/20260521_134807_respuestas_comentario_up.sql")

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
func (m *RespuestasComentario_20260521_134807) Down() {
	 file, err:=ioutil.ReadFile("../scripts/20260521_134807_respuestas_comentario_down.sql")

	if err != nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file),";")
	for _, request:=range requests{
		fmt.Println(requests)
		m.SQL(request)
	}
}
