package router

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/Lazaro-Barros/boilerplate-golang/command/infra/driver/http"
)

type GinRouter struct {
	engine *gin.Engine
}

func NewGinRouter() http.Router {
	return &GinRouter{
		engine: gin.Default(),
	}
}

func (r *GinRouter) GET(path string, handler func(ctx http.Ctx)) {
	r.engine.GET(path, func(c *gin.Context) {
		GinContext := &GinContext{C: c}
		handler(GinContext)
	})
}

func (r *GinRouter) POST(path string, handler func(ctx http.Ctx)) {
	r.engine.POST(path, func(c *gin.Context) {
		GinContext := &GinContext{C: c}
		handler(GinContext)
	})
}

func (r *GinRouter) PUT(path string, handler func(ctx http.Ctx)) {
	r.engine.PUT(path, func(c *gin.Context) {
		GinContext := &GinContext{C: c}
		handler(GinContext)
	})
}

func (r *GinRouter) DELETE(path string, handler func(ctx http.Ctx)) {
	r.engine.DELETE(path, func(c *gin.Context) {
		GinContext := &GinContext{C: c}
		handler(GinContext)
	})
}

func (r *GinRouter) Run(addr string) error {
	return r.engine.Run(addr)
}

type GinContext struct {
	C *gin.Context
}

func (g *GinContext) JSON(code int, obj interface{}) {
	g.C.JSON(code, obj)
}

func (g *GinContext) Bind(obj interface{}) error {
	return g.C.Bind(obj)
}

func (g *GinContext) ShouldBindJSON(obj interface{}) error {
	return g.C.ShouldBindJSON(obj)
}

func (g *GinContext) BindJSON(obj interface{}) error {
	return g.C.BindJSON(obj)
}

func (g *GinContext) Param(key string) string {
	return g.C.Param(key)
}

func (g *GinContext) Query(key string) string {
	return g.C.Query(key)
}

func (g *GinContext) Status(code int) {
	g.C.Status(code)
}

func (g *GinContext) ContextRequest() context.Context {
	return g.C.Request.Context()
}
