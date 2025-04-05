package main

import (
	"fmt"
	"log"
	"net/http"

	"be/handlers"
	"be/utils"
)

func main() {
	// Роутинг
	http.HandleFunc("/double", utils.EnableCORS(handlers.HandleDouble))
	http.HandleFunc("/half", utils.EnableCORS(handlers.HandleHalf))

	// Запуск сервера
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
