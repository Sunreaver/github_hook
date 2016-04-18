package main

import (
	"github.com/go-macaron/pongo2"
	"gopkg.in/macaron.v1"

	"github.com/sunreaver/github_hook/ctl"
)

func main() {
	m := macaron.Classic()
	m.Use(pongo2.Pongoer())
	m.Post("/github.com/sunreaver/:url", ctl.Hook)
	m.Get("/github.com/sunreaver/:url", ctl.Hook)
	m.Run("127.0.0.1", 8097)
}
