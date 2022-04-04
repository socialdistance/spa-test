package sqlstorage

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

const configFile = "config.yaml"

func TestStorage(t *testing.T) {
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		t.Skip(configFile + " file does not exists")
	}

	configContent, _ := os.ReadFile(configFile)

	var config struct {
		Storage struct {
			URL string
		}
	}

	err := yaml.Unmarshal(configContent, &config)
	if err != nil {
		t.Fatal("Failed to unmarshal config", err)
	}

	ctx := context.Background()
	storage := New(ctx, config.Storage.URL)
	if err := storage.Connect(ctx); err != nil {
		t.Fatal("Failed to connect to DB server", err)
	}

	t.Run("test SQL", func(t *testing.T) {
		tx, err := storage.conn.BeginTx(ctx, pgx.TxOptions{
			IsoLevel:       pgx.Serializable,
			AccessMode:     pgx.ReadWrite,
			DeferrableMode: pgx.NotDeferrable,
		})
		if err != nil {
			t.Fatal("Failed to connect to DB server", err)
		}

		_, err = storage.FindAccount("test1")
		if err != nil {
			t.FailNow()
			return
		}
		require.Nil(t, err, nil)

		err = tx.Rollback(ctx)
		if err != nil {
			t.Fatal("Failed to rollback tx", err)
		}
	})
}
