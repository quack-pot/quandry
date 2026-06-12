package internal

import (
	"fmt"
	"log"

	"github.com/quack-pot/quandry/internal/core"
	"github.com/quack-pot/quandry/internal/postgres"
)

func NewDatabase(
	connection_string string,
	driver core.DbDriver,
	options *core.DbOptions,
) (core.IDatabase, error) {
	if options == nil {
		options = core.DefaultDbOptions()
	}

	switch driver {
	case core.DB_DRIVER_POSTGRES:
		return postgres.NewPostgresDatabase(connection_string, options)

	default:
		err := fmt.Errorf("[ERROR]: Invalid database driver provided: %d", driver)
		log.Panic(err.Error())
		return nil, err
	}
}
