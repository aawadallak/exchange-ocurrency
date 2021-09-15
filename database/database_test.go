package database

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection(t *testing.T) {

	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	Init()

	db := GetDatabase()

	err := db.Ping(context.Background())

	assert.NoError(t, err)
}

func TestDatabaseMigratons(t *testing.T) {

	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	Init()

	err := RunMigrations()

	assert.NoError(t, err)
}
