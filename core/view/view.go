package view

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"html/template"
)

//Render function for rendering view tamplate
func Render(ctx *fasthttp.RequestCtx, data interface{}, tmpl ...string) {
	ctx.Response.Header.Set("Content-Type", "text/html")

	tmpls := append([]string{"core/template/root.tpl"}, tmpl...)
	t, err := template.ParseFiles(tmpls...)

	if err != nil {
		fmt.Fprint(ctx, "Error Parsing template: "+err.Error())
	} else {
		err = t.ExecuteTemplate(ctx, "root", data)

		if err != nil {
			fmt.Fprint(ctx, "Error Execute template: "+err.Error())
		}
	}
}
