package repository

import (
	"os"
	"testing"

	"github.com/Daniel-Vinicius/challenge-golang/adapter/repository/fixture"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepository(db)
	err := repository.Insert("1", "1", 12.1)
	assert.Nil(t, err)
}
