package main

import (
	"context"
	"github.com/nomadbala/crust/server/internal/app"
)

var (
	ctx = context.Background()
)

func main() {
	app.Run()
}
