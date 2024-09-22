package interfaces

import (
	"context"
)

type Application interface {
	Run(context.Context)
}

type Computer interface {
	Parse(context.Context, string) ([]string, error)
}

type Storage interface {
	Set(context.Context, string, string) error
	Get(context.Context, string) (string, error)
	Delete(context.Context, string) error
}

type Delivery interface {
	GetRequest(ctx context.Context) (string, error)
	SendResponse(ctx context.Context, msg string, err error) error
}
