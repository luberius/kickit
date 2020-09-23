package server

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func AssetServer(r *router.Router) {
	fs := &fasthttp.FS{
		Root:               "./",
		IndexNames:         []string{"index.html"},
		GenerateIndexPages: true,
		Compress:           false,
		AcceptByteRange:    false,
	}

	fsHandler := fs.NewRequestHandler()

	r.GET("/asset/{any:*}", fsHandler)
}
