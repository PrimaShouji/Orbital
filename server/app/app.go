package app

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var Store *sessions.FilesystemStore

type App struct {
	*AppPath
	engine *gin.Engine
	port   uint64
}

type AppPath struct {
	router *gin.RouterGroup
}

func Init(port uint64) *App {
	Store = sessions.NewFilesystemStore("", []byte(os.Getenv("PRIMA_ORBITAL_AUTH0_SECRET")))
	gob.Register(map[string]interface{}{})

	r := gin.Default()

	return &App{
		AppPath: &AppPath{
			router: &r.RouterGroup,
		},
		engine: r,
		port:   port,
	}
}

func (a *AppPath) Get(relativePath string, handler func(*gin.Context)) {
	a.router.GET(relativePath, handler)
}

func (a *AppPath) Group(relativePath string, handler func(*gin.Context)) *AppPath {
	return &AppPath{
		router: a.router.Group(relativePath, handler),
	}
}

func (a *App) Run() {
	a.engine.Run(":" + fmt.Sprint(a.port))
}
