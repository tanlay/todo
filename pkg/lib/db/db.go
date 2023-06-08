/*
配置数据库连接
*/

package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"todo/config"
)

var (
	GlobalDB *gorm.DB
)

var (
	db  *gorm.DB
	err error
)

func GlobalWithDSN(conf config.DatabaseConf) (*gorm.DB, error) {
	mysqlPrefix := "mysql://"
	if strings.HasPrefix(conf.DSN, mysqlPrefix) {
		db, err = gorm.Open(mysql.Open(strings.ReplaceAll(conf.DSN, mysqlPrefix, "")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error), //默认应该是Warn，设置Error则不显示慢SQL
		})
		if err != nil {
			return nil, err
		}

		sqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}
		if err := sqlDB.Ping(); err != nil {
			return nil, err
		}
		//配置连接池 TODO

		return db, nil
	} else {
		return nil, gorm.ErrUnsupportedDriver
	}
}

func GlobalWithDB(conf config.DatabaseConf) error {
	db, err := GlobalWithDSN(conf)
	if err != nil {
		return err
	}
	GlobalDB = db
	return nil
}
