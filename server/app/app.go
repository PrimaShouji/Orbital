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
	router *gin.Engine
	port   uint64
}

func Init(port uint64) *App {
	Store = sessions.NewFilesystemStore("", []byte(os.Getenv("PRIMA_ORBITAL_AUTH0_SECRET")))
	gob.Register(map[string]interface{}{})

	r := gin.Default()

	return &App{
		router: r,
		port:   port,
	}
}

func (a *App) Get(relativePath string, handler func(*gin.Context)) {
	a.router.GET(relativePath, handler)
}

func (a *App) Run() {
	a.router.Run(":" + fmt.Sprint(a.port))
}
