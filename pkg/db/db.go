package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"adv-demo/configs"
)

type Db struct {
	db *gorm.DB
}

func New(config *configs.Config) *Db {
	// Открываем соединение передаем DSN + указатель на конфиг, если нужно будет что то модифичировать в дефолтном подключении
	db, err := gorm.Open(postgres.Open(config.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("INFO: Соединение с БД установлено")
	return &Db{
		db: db,
	}
}
