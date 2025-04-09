package main

import (
	"fmt"
	"log"
	"net/http"

	"be/handlers"
	"be/utils"
)

func main() {
	http.HandleFunc("/save-message", utils.EnableCORS(handlers.HandleSaveMessage))

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
