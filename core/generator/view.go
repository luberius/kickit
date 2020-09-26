package generator

import (
	"bytes"
	"fmt"
	"kickit/core/data"
	"strings"
	"text/template"
)

func GenerateView(table *data.Table) error {
	var viewTpl bytes.Buffer

	t, err := template.ParseFiles("template/create.kicktpl")
	if err != nil {
		return err
	}

	if err := t.Execute(&viewTpl, table); err != nil {
		return err
	}

	viewPath := fmt.Sprintf("resources\\views\\%s\\", strings.ToLower(table.TableName))
	project := data.GetProject()

	err = CreateFile(viewTpl.String(), project.RootPath + viewPath, fmt.Sprintf("%s.php", "create.blade"))
	return err
}