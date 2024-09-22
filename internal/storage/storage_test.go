package storage_test

import (
	"context"
	"errors"
	"testing"

	"github.com/skyris/KVlight/internal/storage"
	"github.com/skyris/KVlight/pkg/issues"
)

func TestSimpleStore(t *testing.T) {
	ctx := context.Background()
	testKey := "foo"
	nonExistentKey := "nonexistent"
	expectedValue := "bar"

	store := storage.NewSimpleStore()

	// Test Set and Get
	err := store.Set(ctx, testKey, expectedValue)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	gotValue, err := store.Get(ctx, testKey)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if gotValue != expectedValue {
		t.Errorf("expected %s, got '%s'", expectedValue, gotValue)
	}

	// Test Get with non-existent key
	_, err = store.Get(ctx, nonExistentKey)
	if !errors.Is(err, issues.ErrInvalidKey) {
		t.Errorf("expected ErrInvalidKey, got %v", err)
	}

	// Test Delete with existent key
	err = store.Delete(ctx, testKey)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	// Test Delete with non-existent key
	err = store.Delete(ctx, nonExistentKey)
	if !errors.Is(err, issues.ErrInvalidKey) {
		t.Errorf("expected ErrInvalidKey, got %v", err)
	}

	// Test Get after Delete
	_, err = store.Get(ctx, testKey)
	if !errors.Is(err, issues.ErrInvalidKey) {
		t.Errorf("expected ErrInvalidKey after delete, got %v", err)
	}
}
