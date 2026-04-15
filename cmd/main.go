package main

import (
	"adv-demo/configs"
	"adv-demo/internal/auth"
	"adv-demo/internal/store"

	"context"
	"os/signal"
	"syscall"
	"time"

	"fmt"
	"net/http"
)

func main() {
	// Инциируем контекст с сигналами
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer stop()

	// Инициируем модули приложения
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	db := store.New(conf.Db.FilePath)
	db.Load()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
		Store:  db,
	})

	// Подготавливаем и запускаем HTTP-сервер
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	go func() {
		fmt.Println("Server is listining on port 8081")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println("ERROR: Запуск сервера")
		}
	}()

	// Graceful shutdown
	// Ждем сигнал завершения
	<-ctx.Done()
	fmt.Println("INFO: Получен сигнал для остановки приложения")
	// Время на завершение всех операций
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Останавливаем HTTP сервер
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Println("ERROR: Ошибка при завершении работы HTTP-сервера:", err)
	} else {
		fmt.Println("INFO: ДаHTTP-сервера успешно остановлен")
	}

	// Сохраняем данные на диск
	if err := db.Snapshoot(); err != nil {
		fmt.Println("ERROR: Ошибка при сохранение данных на диск", err)
	} else {
		fmt.Println("INFO: Данные успешно сохранены на диск")
	}
	fmt.Println("INFO: Приложение остановлено")

}
