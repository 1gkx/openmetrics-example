package main

import (
	server "github.com/1gkx/openmetrics/internal"
)

func main() {
	if err := server.Start("8180"); err != nil {
		panic(err)
	}
}

