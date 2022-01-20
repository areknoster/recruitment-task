package storage

import (
	"context"
	"errors"
)

var (
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrCtxDone           = errors.New("context done")
)

type ImportantData struct {
	Message string
}

// Storage let's you retrieve some data if you pass correct password. It's safe for concurrent use.
type Storage interface {
	Get(ctx context.Context, key string, password string) (ImportantData, error)
}
