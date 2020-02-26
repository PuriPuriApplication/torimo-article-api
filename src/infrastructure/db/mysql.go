package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"
)

type env struct {
	DB_User  string    `default:"torimo"`
	DB_Password string `default:"torimo"`
	DB_Dbname string   `default:"torimo"`
	DB_Url string      `default:"127.0.0.1:3306"`
}

func Init() *gorm.DB {
	// Set environment
	var env env
	envconfig.Process("", &env)

	// Connect to mysql
	connect := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", env.DB_User, env.DB_Password, env.DB_Url, env.DB_Dbname)
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		panic("failed to connect database")
	}

	// Log setting
	db.LogMode(true)

	return db
}
