package generator

import (
	"bytes"
	"fmt"
	"kickit/core/data"
	"text/template"
)

func GenerateModel(model *data.Table) error {
	var modelTpl bytes.Buffer
	t, err := template.ParseFiles("template/model.kicktpl")
	if err != nil {
		return err
	}

	if err := t.Execute(&modelTpl, model); err != nil {
		return err
	}

	err = CreateFile(modelTpl.String(), fmt.Sprintf("%s%s.php", "", model.Name))
	return err
}
