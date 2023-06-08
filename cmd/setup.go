package cmd

import (
	"todo/config"
	"todo/pkg/lib/db"
)

func SetupGlobalDB(config config.Config) error {
	return db.GlobalWithDB(config.Database)
}
