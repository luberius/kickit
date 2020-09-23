package main

import (
	"kickit/core/data"
	"kickit/core/generator"
	"strings"
)

func main() {
	m := data.Model{
		TableName: "profile",
		Fields: []data.Field{
			{
				Name:      "fullname",
				FieldType: "varchar",
			},
			{
				Name:      "age",
				FieldType: "integer",
			},
		},
	}

	m.Name = strings.Title(m.TableName)

	err := generator.GenerateModel(m)
	if err != nil {
		panic(err)
	}
}
