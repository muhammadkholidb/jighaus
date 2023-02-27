package main

import (
	"embed"
	"jighaus/cmd"
)

//go:embed public/*
var public embed.FS

func main() {
	cmd.Server(public)
}
