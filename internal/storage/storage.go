package storage

import (
	"context"

	"github.com/skyris/KVlight/pkg/issues"
)

type Storage interface {
	Set(context.Context, string, string) error
	Get(context.Context, string) (string, error)
	Delete(context.Context, string) error
}

type SimpleEngine struct {
	db map[string]string
}

var _ Storage = &SimpleEngine{}

func (d *SimpleEngine) Set(_ context.Context, key, value string) error {
	d.db[key] = value
	return nil
}

func (d *SimpleEngine) Get(_ context.Context, key string) (string, error) {
	value, ok := d.db[key]
	if !ok {
		return "", issues.ErrInvalidKey
	}
	return value, nil
}

func (d *SimpleEngine) Delete(_ context.Context, key string) error {
	if _, ok := d.db[key]; !ok {
		return issues.ErrInvalidKey
	}
	delete(d.db, key)
	return nil
}

func NewSimpleStore() *SimpleEngine {
	return &SimpleEngine{
		db: make(map[string]string),
	}
}
