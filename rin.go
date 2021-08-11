package rin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Rin struct {
	*gin.Engine
}

// New creates a Rin application instance
func New() *Rin {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		MaxAge: 50 * time.Second,
		AllowCredentials: true,
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