package data

type Field struct {
	Name      string
	FieldType string
}

type Table struct {
	Name      string
	TableName string
	Fields    []Field
}
