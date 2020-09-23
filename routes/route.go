package routes

import "github.com/fasthttp/router"
import "kickit/controllers/dashboard"

func Handle(r *router.Router) {
	r.GET("/", dashboard.Index)
}
