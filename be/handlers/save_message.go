package handlers

import (
	"encoding/json"
	"net/http"

	"be/service"
	"be/utils"
)

var msgService *service.MessageService

func InitHandlers(service *service.MessageService) {
	msgService = service
}

func HandleSaveMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendJSONError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	// Вызываем сервис с контекстом
	savedMsg, err := msgService.SaveMessage(ctx, req.Message)
	if err != nil {
		utils.SendJSONError(w, "Ошибка сохранения", http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, map[string]string{"saved": savedMsg.Text})
}
