package main

import (
	"context"

	"github.com/skyris/KVlight/internal/compute"
	"github.com/skyris/KVlight/internal/database"
	"github.com/skyris/KVlight/internal/delivery"
	"github.com/skyris/KVlight/internal/storage"
)

func main() {
	ctx := context.Background()
	db := database.NewDataBase(
		compute.NewCompute(),
		storage.NewSimpleStore(),
		delivery.Default(),
	)
	db.Run(ctx)
}
