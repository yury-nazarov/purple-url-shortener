# URL Shortener

## CLI
- `go mod init <name>` - инициализация go module
- `go get <dependency>` - добавляет зависимость в `go.mod`, скачивает пакет в кеш `~/go/pkg/mod`
    - `Зависимости / dependency` - В Go зависимости не устанавливаются в проект, они добавляются в `go.mod` проекта и скачиваются в кеш `~/go/pkg/mod`
                Если нужно удалить из проекта пакет 
                - Удаляем все импорты с ним
                - go mod tidy - сам поправит `go.mod`, `go.sum`
    - `go.mod` - список зависимостей конкретного проекта
    - `go.sum` - контрольные сумму всех зависимостей, гарантирует, что ты скачал тот же код, а не подмененный.
- `go mod tidy` актуализирует `go.mod`, `go.sum` 
    - Удаляет из `go.mod`, `go.sum` зависимостей которых больше нет в коде (import удален)
    - Добавляет в `go.mod`, `go.sum` зависимости которые появились в коде (import добавлен)
- `go clean -modcache` - удаляет все пакеты из кеша


## Прочее
- [Выбор ORM](https://github.com/d-tsuji/awesome-go-orms) - ORM(Object Relation Mapping) способ работать с БД через объекты в коде, а не через SQL на прямую.
- [gorm.io](gorm.io) - одна из самых популярных ORM в Go
- [https://atlasgo.io/](https://atlasgo.io/) - инструмент для управления схемой БД (Shchema as a Code)
    - 👉 GORM — это runtime слой (ORM)
    - 👉 Atlas — это build/deploy слой (schema management)