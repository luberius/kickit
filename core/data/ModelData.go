package data

type Field struct {
	Name string
	FieldType string
}

type Model struct {
	Name   string
	TableName   string
	Fields []Field
}
