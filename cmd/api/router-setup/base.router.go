package routersetup

import httpserver "http101/cmd/api/http-server"

type RouterMap struct {
	routers *httpserver.Router
}

func New(routers *httpserver.Router) *RouterMap {
	return &RouterMap{routers}
}
