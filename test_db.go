package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

type Table struct {
	name     string
	dataType string
}

func main() {
	db, err := sql.Open("mysql", "root:@(localhost:3306)/dizifood")
	err = db.Ping()

	results, err := db.Query("SHOW COLUMNS FROM invoice")

	defer results.Close()

	fmt.Println("=== Show Table ===")

	var tables []Table

	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {
		var table Table
		err := results.Scan(&table.name, &table.dataType)

		if err != nil {
			fmt.Println(err.Error())
			break
		}

		tables = append(tables, table)
	}

	for _, t := range tables {
		fmt.Println(t.name)
	}
}
