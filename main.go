package main

import (
	"log"
	"net/http"
	"password-manager/api"
	"password-manager/config"
	"password-manager/db"

	"github.com/gorilla/mux"
)

func main() {
	// Инициализация базы данных
	database, err := db.InitDB("passwords.db")
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer database.Close()

	repo := db.NewRepository(database)

	apiHandler := &api.API{Repo: repo}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Password Manager API is running!"}`))
	}).Methods("GET")

	router.HandleFunc("/generate", apiHandler.GeneratePasswordHandler) // УТОЧНИТЬ!
	router.HandleFunc("/search", apiHandler.SearchPasswordsHandler).Methods("GET")

	router.HandleFunc("/passwords", apiHandler.CreatePasswordHandler).Methods("POST")
	router.HandleFunc("/passwords", apiHandler.GetPasswordsHandler).Methods("GET")
	router.HandleFunc("/passwords/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"id": "` + id + `"}`))
	}).Methods("GET")

	log.Println("Сервер запущен на http://localhost" + config.ServerPort)
	if err := http.ListenAndServe(config.ServerPort, router); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
