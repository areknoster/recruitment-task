package storage

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"time"
)

const (
	passwordHashHex = "456aab964fc7850968b3ce8be6be211abe741ca6f8396ed71e4dfea903e3a11c66dc52867b022edd647908bddbd87ef4181c2cdc4c3a7b6ab6b7ab433b300a1b"
	hashLen         = 64
)

var storedPasswordHash []byte

func init() {
	var err error
	storedPasswordHash, err = hex.DecodeString(passwordHashHex)
	if err != nil || len(storedPasswordHash) != hashLen {
		panic("worked on my machine :(")
	}
	rand.Seed(time.Now().Unix())
}

type ImportantStorage struct{}

func (ids ImportantStorage) Get(ctx context.Context, key string, password string) (ImportantData, error) {
	sleepTime := time.Duration(rand.Intn(10)) * time.Second
	passwordHash := sha512.Sum512([]byte(password))
	select {
	case <-ctx.Done():
		return ImportantData{}, ErrCtxDone
	case <-time.After(sleepTime):
		if !bytes.Equal(passwordHash[:], storedPasswordHash[:]) {
			return ImportantData{}, ErrIncorrectPassword
		}

		return ImportantData{
			Message: "I slept for " + sleepTime.String(),
		}, nil
	}
}
