package db

import (
	"database/sql"
	"password-manager/models"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreatePassword(p *models.Password) error {
	query := "INSERT INTO passwords (service_name, username, password) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, p.ServiceName, p.Username, p.Password)
	return err
}

func (r *Repository) GetPasswords() ([]models.Password, error) {
	query := "SELECT id, service_name, username, password, created_at FROM passwords"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passwords []models.Password
	for rows.Next() {
		var p models.Password
		if err := rows.Scan(&p.ID, &p.ServiceName, &p.Username, &p.Password, &p.CreatedAt); err != nil {
			return nil, err
		}
		passwords = append(passwords, p)
	}
	return passwords, nil
}

func (r *Repository) SearchPasswords(serviceName string) ([]models.Password, error) {
	query := "SELECT id, service_name, username, password, created_at FROM passwords WHERE service_name LIKE ?"
	rows, err := r.DB.Query(query, "%"+serviceName+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passwords []models.Password
	for rows.Next() {
		var p models.Password
		if err := rows.Scan(&p.ID, &p.ServiceName, &p.Username, &p.Password, &p.CreatedAt); err != nil {
			return nil, err
		}
		passwords = append(passwords, p)
	}
	return passwords, nil
}
