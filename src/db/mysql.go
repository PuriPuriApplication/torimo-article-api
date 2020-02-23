package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"
)

type env struct {
	User  string    `default:"torimo"`
	Password string `default:"torimo"`
	Dbname string   `default:"torimo"`
	Url string      `default:"127.0.0.1:3306"`
}

func Init() *gorm.DB {
	// Set environment
	var env env
	envconfig.Process("", &env)

	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", env.User, env.Password, env.Url, env.Dbname),
	)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
