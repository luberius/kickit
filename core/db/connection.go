package db

import (
	"database/sql"
	"fmt"
	"kickit/core/data"
)
import _ "github.com/go-sql-driver/mysql"

func Mysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@(localhost:3306)/dizifood")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Connection() (*sql.DB, error) {
	return Mysql()
}

func GetTable(name string) (*data.Table, error) {
	db, err := Connection()
	if err != nil {
		fmt.Println(err)
	}

	results, err := db.Query(fmt.Sprintf("SELECT `COLUMN_NAME`, `DATA_TYPE` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`='dizifood' AND `TABLE_NAME`='%s';", name))

	if err != nil {
		return nil, err
	}

	var fields []data.Field
	for results.Next() {
		var field data.Field
		err := results.Scan(&field.Name, &field.FieldType)

		if err != nil {
			fmt.Println(err.Error())
			break
		}

		fields = append(fields, field)
	}

	return &data.Table{
		TableName: name,
		Fields:    fields,
	}, nil
}
