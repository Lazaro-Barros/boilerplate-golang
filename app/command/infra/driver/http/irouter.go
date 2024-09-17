package http

import "context"

type Router interface {
	Run(addr string) error
	GET(path string, handler func(ctx Ctx))
	POST(path string, handler func(ctx Ctx))
	PUT(path string, handler func(ctx Ctx))
	DELETE(path string, handler func(ctx Ctx))
}

type Ctx interface {
	JSON(code int, obj interface{})
	Bind(obj interface{}) error
	ShouldBindJSON(obj interface{}) error
	BindJSON(obj interface{}) error
	Param(key string) string
	Query(key string) string
	Status(code int)
	ContextRequest() context.Context
}
