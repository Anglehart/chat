package handlers

import (
	"encoding/json"
	"net/http"

	"be/service"
	"be/utils"
)

func HandleSaveMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendJSONError(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendJSONError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	result := service.SaveMessage(req.Message)
	utils.SendJSONResponse(w, map[string]string{"result": result})
}
