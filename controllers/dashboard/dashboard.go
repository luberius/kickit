package dashboard

import (
	"github.com/valyala/fasthttp"
	"kickit/core/view"
)

func Index(ctx *fasthttp.RequestCtx) {
	view.Render(ctx, "World", "views/dashboard.tpl", "views/layout/header.tpl", "views/layout/footer.tpl")
}
