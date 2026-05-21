package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Nutricion_20260521_134738 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Nutricion_20260521_134738{}
	m.Created = "20260521_134738"

	migration.Register("Nutricion_20260521_134738", m)
}

// Run the migrations
func (m *Nutricion_20260521_134738) Up() {
	 file, err:=ioutil.ReadFile("../scripts/20260521_134738_nutricion_up.sql")

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
func (m *Nutricion_20260521_134738) Down() {
	 file, err:=ioutil.ReadFile("../scripts/20260521_134738_nutricion_down.sql")

	if err != nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file),";")
	for _, request:=range requests{
		fmt.Println(requests)
		m.SQL(request)
	}
}

