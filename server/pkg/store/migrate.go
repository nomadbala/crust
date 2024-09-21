package store

import (
	"errors"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(dataSourceName string) error {
	if !strings.Contains(dataSourceName, "://") {
		return errors.New("store: undefined data source name " + dataSourceName)
	}

	migrations, err := migrate.New(fmt.Sprintf("file://db/postgres/migration"), dataSourceName)
	if err != nil {
		return err
	}

	if err = migrations.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return err
	}

	return nil
}
