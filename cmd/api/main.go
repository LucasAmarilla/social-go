package main

import (
	"log"

	"github.com/lucasamarilla/social-go/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":42069"),
	}
	app := &app{
		config: cfg,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
