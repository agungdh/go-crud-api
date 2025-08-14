package main

import (
	"log"
	"os"

	"github.com/agungdh/go-crud-api/router"
)

type AppDeps struct {
	// contoh dependencies: logger, db, config, dll
	Logger *log.Logger
}

func main() {
	deps := &AppDeps{
		Logger: log.New(os.Stdout, "[myapp] ", log.LstdFlags|log.Lshortfile),
	}

	r := router.New(&router.Deps{
		Logger: deps.Logger, // cocok dengan interface Printf di router.Deps
	})
	
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
