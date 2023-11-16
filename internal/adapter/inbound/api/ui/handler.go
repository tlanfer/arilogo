package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed app/dist/**
var app embed.FS

func NewHandler() http.Handler {
	sub, _ := fs.Sub(app, "app/dist")
	return http.FileServer(http.FS(sub))
}
