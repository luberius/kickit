package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type Field struct {
	Name      string
	FieldType string
}

type ModelData struct {
	Name      string
	TableName string
	Fields    []Field
}

func createModelData() ModelData {
	m := ModelData{
		TableName: "profile",
		Fields: []Field{
			{Name: "fullname", FieldType: "varchar"},
			{Name: "age", FieldType: "number"},
		},
	}
	m.Name = strings.Title(m.TableName)
	return m
}

func main() {
	var modelTpl bytes.Buffer
	t, err := template.ParseFiles("template/model.kicktpl")
	if err != nil {
		panic(err)
	}

	if err := t.Execute(&modelTpl, createModelData()); err != nil {
		panic(err)
	}

	fmt.Print(modelTpl.String())
}
