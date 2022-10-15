package controller

import (
	"encoding/json"
	"live-tests/internal/infra/database"
	"live-tests/internal/usecase"
	"net/http"
)

func (b *BaseHandler) CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	var content usecase.CreateClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		w.WriteHeader((http.StatusBadRequest))
		return
	}
	repository := database.NewClientRepository(b.db)
	uc := usecase.NewClientUseCase(repository)
	_, err = uc.Execute(content)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
