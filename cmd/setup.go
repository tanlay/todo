package cmd

import (
	"github.com/tanlay/todo/config"
	"github.com/tanlay/todo/pkg/lib/db"
)

func SetupGlobalDB(config config.Config) error {
	return db.GlobalWithDB(config.Database)
}
