package web_api

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed static/*
var uiFiles embed.FS

func registerUI(engine *gin.Engine) {
	sub, err := fs.Sub(uiFiles, "static")
	if err != nil {
		return
	}
	engine.StaticFS("/ui", http.FS(sub))
}
