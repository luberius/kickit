package main

import (
	"fmt"
	"kickit/core/db"
	"kickit/core/generator"
	"strings"
)

func main() {
	m, err := db.GetTable("invoice")

	if err != nil {
		panic(err)
	}

	m.Name = strings.Title(m.TableName)

	fmt.Println(m)

	err = generator.GenerateModel(m)
	if err != nil {
		panic(err)
	}
}
