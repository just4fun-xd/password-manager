package api

import (
	"encoding/json"
	"net/http"
	"password-manager/db"
	"password-manager/models"
	"password-manager/utils"
)

type API struct {
	Repo *db.Repository
}

type GeneratePasswordRequest struct {
	Length     int  `json:"length"`
	UseDigits  bool `json:"use_digits"`
	UseSymbols bool `json:"use_symbols"`
}

type GeneratePasswordResponse struct {
	Password string `json:"password"`
}

func (api *API) GeneratePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var req GeneratePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	password := utils.GeneratePassword(req.Length, req.UseDigits, req.UseSymbols)
	res := GeneratePasswordResponse{Password: password}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (api *API) CreatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Password
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := api.Repo.CreatePassword(&p); err != nil {
		http.Error(w, "Failed to create password", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (api *API) GetPasswordsHandler(w http.ResponseWriter, r *http.Request) {
	passwords, err := api.Repo.GetPasswords()
	if err != nil {
		http.Error(w, "Failed to fetch passwords", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(passwords)
}

func (api *API) SearchPasswordsHandler(w http.ResponseWriter, r *http.Request) {
	serviceName := r.URL.Query().Get("service")
	passwords, err := api.Repo.SearchPasswords(serviceName)
	if err != nil {
		http.Error(w, "Failed to search passwords", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(passwords)
}
