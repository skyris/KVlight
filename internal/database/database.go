package database

import (
	"context"
	"errors"
	"io"
	"log"

	"github.com/skyris/KVlight/pkg/interfaces"
	"github.com/skyris/KVlight/pkg/issues"
	"github.com/skyris/KVlight/pkg/types"
)

type App struct {
	computer interfaces.Computer
	storage  interfaces.Storage
	delivery interfaces.Delivery
}

func NewDataBase(
	computer interfaces.Computer,
	storage interfaces.Storage,
	delivery interfaces.Delivery,
) *App {
	return &App{
		computer: computer,
		storage:  storage,
		delivery: delivery,
	}
}

func (a *App) Run(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("crushing caused by:", err)
		}
		log.Println("shutting down...")
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("End of Loop.")
			return
		default:
			line, err := a.delivery.GetRequest(ctx)
			if err != nil {
				if errors.Is(err, io.EOF) {
					log.Println("Got CTRL+D from user")
					return
				}
				log.Printf("Message recieving error: %v\n", err)
				break
			}

			args, err := a.computer.Parse(ctx, line)
			if err != nil {
				log.Printf("Parsing error: %v\n", err)
				continue
			}

			output, err := a.dispatch(ctx, args)
			err = a.delivery.SendResponse(ctx, output, err)
			if err != nil {
				log.Println("Handling error:", err)
				continue
			}
		}
	}
}

func (a *App) dispatch(ctx context.Context, args []string) (string, error) {
	switch args[0] {
	case types.CommandSET:
		return "", a.storage.Set(ctx, args[1], args[2])
	case types.CommandGET:
		return a.storage.Get(ctx, args[1])
	case types.CommandDEL:
		return "", a.storage.Delete(ctx, args[1])
	}
	return "", issues.ErrInvalidCommand
}
