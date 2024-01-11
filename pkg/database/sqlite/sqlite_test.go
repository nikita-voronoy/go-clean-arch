package sqlite

import (
	"context"
	"github.com/google/uuid"
	"github.com/nikita-voronoy/go-clean-arch/config"
	"github.com/nikita-voronoy/go-clean-arch/internal/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
	cfg := &config.Config{
		DB: config.DatabaseConfig{
			Sqlite: config.SqliteConfig{
				DatabasePath: ":memory:",
			},
		},
	}

	db, err := NewDatabase(cfg)
	assert.NoError(t, err, "Failed to create new database")

	assert.NotNil(t, db, "Database is nil")
}

func TestCreate(t *testing.T) {
	cfg := &config.Config{
		DB: config.DatabaseConfig{
			Sqlite: config.SqliteConfig{
				DatabasePath: ":memory:",
			},
		},
	}

	db, err := NewDatabase(cfg)
	assert.NoError(t, err, "Failed to create new database")

	user := &entities.User{
		ID:       uuid.New(),
		Username: "user",
		Password: "pass",
		Email:    "email@gmail.com",
		Metadata: entities.Metadata{
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			LastLoginAt: time.Now(),
		},
		Token: "token",
	}

	err = db.Create(context.Background(), user)
	assert.NoError(t, err, "Failed to create user")
}
