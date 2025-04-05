package handlers

import (
	"encoding/json"
	"net/http"

	"be/service"
	"be/utils"
)

func HandleDouble(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendJSONError(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Number int `json:"number"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendJSONError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	result := service.Double(req.Number)
	utils.SendJSONResponse(w, map[string]int{"result": result})
}
