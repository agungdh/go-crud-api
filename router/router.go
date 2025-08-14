package router

import (
	"github.com/agungdh/go-crud-api/handler"
	"github.com/agungdh/go-crud-api/middleware"
	"github.com/gin-gonic/gin"
)

type Deps struct {
	Logger interface {
		Printf(format string, v ...any)
	}
}

func New(d *Deps) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())           // log access
	r.Use(middleware.RequestID()) // contoh middleware custom

	// Health & meta
	r.GET("/health", handler.HealthHandler())

	r.GET("/project", handler.ProjectHandler())

	return r
}
