package app

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"kickit/core/server"
	"kickit/routes"
)

func Run(addr string) {
	r := router.New()
	server.AssetServer(r)
	routes.Handle(r)

	fmt.Printf("Running Kickit. on %s", addr)

	panic(fasthttp.ListenAndServe(addr, r.Handler))
}
