package main

import (
	"context"
	"log"
	"os"

	"github.com/ShayeGun/go-server/internal/common"
	"github.com/ShayeGun/go-server/internal/service"
	db "github.com/ShayeGun/go-server/internal/storage/postgres"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Warning: No .env file found")
		}
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

	services, _ := service.NewService(
		common.ExternalDependencies{
			RepositoryInterface: store,
		},
	)

	mux := app.mount(services)
	app.run(mux)

	// TODO: close repository connection
	// defer conn.Close(ctx)
}
