package generator

import (
	"bytes"
	"fmt"
	"kickit/core/data"
	"text/template"
)

func GenerateModel(table *data.Table) error {
	var modelTpl bytes.Buffer
	t, err := template.ParseFiles("template/model.kicktpl")
	if err != nil {
		return err
	}

	if err := t.Execute(&modelTpl, table); err != nil {
		return err
	}

	project := data.GetProject()

	err = CreateFile(modelTpl.String(), project.RootPath + "app\\", fmt.Sprintf("%s.php", table.Name))
	return err
}
