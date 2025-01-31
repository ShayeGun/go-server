package main

import (
	"context"
	"log"
	"os"

	"github.com/ShayeGun/go-server/internal/routes"
	db "github.com/ShayeGun/go-server/internal/storage/postgres"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: os.Getenv("PORT"),
	}

	app := &application{
		config: cfg,
	}

	store, err := db.NewRepository(ctx, os.Getenv("DB_URI"))

	if err != nil {
		panic(err)
	}

	ed := &routes.ExternalDependencies{
		RepositoryInterface: store,
	}

	mux := app.mount(ed)
	app.run(mux)

	// TODO: close repository connection
	// defer conn.Close(ctx)
}
