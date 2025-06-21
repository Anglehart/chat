package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"be/handlers"
	"be/service"
	"be/storage"
	"be/utils"
)

func main() {
	// 1. Параметры подключения к PostgreSQL
	connStr := "user=postgres dbname=postgres password=admin host=localhost port=5432 sslmode=disable"

	// 2. Инициализация хранилища с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbStorage, err := storage.NewPostgresStorage(ctx, connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
		os.Exit(1)
	}
	defer dbStorage.Close() // Закрываем соединение при завершении

	// 2. Инициализация сервиса
	msgService := service.NewMessageService(dbStorage)

	// 3. Инициализация обработчиков (передаём сервис)
	handlers.InitHandlers(msgService)

	// 4. Настройка роутинга
	http.HandleFunc("/save-message", utils.EnableCORS(handlers.HandleSaveMessage))
	http.HandleFunc("/get-messages", utils.EnableCORS(handlers.HandleGetMessages))

	// 5. Запуск сервера
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
