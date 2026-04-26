package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"adv-demo/configs"
)

type Db struct {
	DB *gorm.DB
}

func New(config *configs.Config) *Db {
	// Открываем соединение передаем DSN + указатель на конфиг,
	// если нужно будет что то модифицировать в дефолтном подключении
	db, err := gorm.Open(postgres.Open(config.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("INFO: Соединение с БД установлено")
	return &Db{
		DB: db,
	}
}
