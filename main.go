package main

import (
	"github.com/harsht283raj/gol/server"
)

func main() {
	a := server.App{}
	a.Initialize(
		"postgres",
		"admin",
		"gol")

	a.Run(":8010")
}
