package rin

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"time"
)

type Rin struct {
	*gin.Engine
}

// New creates a Rin application instance
func New() *Rin {
	engine := gin.Default()
	engine.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))

	return &Rin{Engine: engine}
}

// Controller Register Controller to Rin application
func (a *Rin) Controller(c IController) {
	g := a.Group(c.GetName())
	c.Register(a, g)
}

// ResponseHandler gin.Context handler with IResponse
type ResponseHandler func(*gin.Context) IResponse