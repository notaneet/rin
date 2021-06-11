package rin

import "github.com/gin-gonic/gin"

type IController interface {
	GetName() string
	Register(app *Rin, group *gin.RouterGroup)

	GET(path string, handlers ...ResponseHandler) IController
	POST(path string, handlers ...ResponseHandler) IController
	PUT(path string, handlers ...ResponseHandler) IController
	DELETE(path string, handlers ...ResponseHandler) IController

	SUB(path string) IController
}

type controllerImpl struct {
	name string

	post map[string][]ResponseHandler
	get  map[string][]ResponseHandler
	put  map[string][]ResponseHandler
	delete map[string][]ResponseHandler

	subControllers []*subControllerImpl
}

type subControllerImpl struct {
	IController
}

func Controller(name string) IController {
	c := &controllerImpl{name: name}
	c.get = map[string][]ResponseHandler{}
	c.post = map[string][]ResponseHandler{}
	c.put = map[string][]ResponseHandler{}
	c.delete = map[string][]ResponseHandler{}

	return c
}

func (c controllerImpl) GetName() string {
	return c.name
}

func (c controllerImpl) Register(app *Rin, group *gin.RouterGroup) {
	for _, controller := range c.subControllers {
		controller.Register(app, group.Group(controller.GetName()))
	}

	for path, handlers := range c.get {
		group.GET(path, rin2gin(handlers)...)
	}

	for path, handlers := range c.post {
		group.POST(path, rin2gin(handlers)...)
	}

	for path, handlers := range c.put {
		group.PUT(path, rin2gin(handlers)...)
	}

	for path, handlers := range c.delete {
		group.DELETE(path, rin2gin(handlers)...)
	}
}

func (c *controllerImpl) GET(path string, handlers ...ResponseHandler) IController {
	for _, handler := range handlers {
		c.get[path] = append(c.get[path], handler)
	}

	return c
}

func (c *controllerImpl) POST(path string, handlers ...ResponseHandler) IController {
	for _, handler := range handlers {
		c.post[path] = append(c.post[path], handler)
	}

	return c
}

func (c *controllerImpl) PUT(path string, handlers ...ResponseHandler) IController {
	for _, handler := range handlers {
		c.put[path] = append(c.put[path], handler)
	}

	return c
}

func (c *controllerImpl) DELETE(path string, handlers ...ResponseHandler) IController {
	for _, handler := range handlers {
		c.delete[path] = append(c.delete[path], handler)
	}

	return c
}

// SUB creates a subController
func (c *controllerImpl) SUB(path string) IController {
	controller := &subControllerImpl{Controller(path)}
	c.subControllers = append(c.subControllers, controller)
	return controller
}

func rin2gin(handlers []ResponseHandler) (ginHandlers []gin.HandlerFunc) {
	for _, handler := range handlers {
		ginHandlers = append(ginHandlers, func(ctx *gin.Context) {
			resp := handler(ctx)
			ctx.JSON(resp.GetStatusCode(), resp.GetResponse())
		})
	}

	return ginHandlers
}