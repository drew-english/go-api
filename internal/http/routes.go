package http

import (
	"github.com/gin-gonic/gin"
)

type requestMethod string
const (
	get requestMethod = "GET"
	post requestMethod = "POST"
	patch requestMethod = "PATCH"
	delete requestMethod = "DELETE"
)

type route struct {
	method requestMethod
	path string
	handler gin.HandlerFunc
}

var v1Routes = []route {
	{ post, "/app_setting", func(c *gin.Context) {} },
	{ get, "/app_setting/:name", func(c *gin.Context) {} },
	{ patch, "/app_setting/:name", func(c *gin.Context) {} },
	{ delete, "/app_setting/:name", func(c *gin.Context) {} },
}

func ConfigRoutes(engine *gin.Engine) {
	apiV1 := engine.Group("/api/v1")
	
	for _, r := range v1Routes {
		apiV1.Handle(string(r.method), r.path, r.handler)
	}
}