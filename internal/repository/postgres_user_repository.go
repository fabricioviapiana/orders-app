package repository

import (
	"database/sql"

	"github.com/fabricioviapiana/orders-app/internal/domain"
)

type postgresUserRespository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *postgresUserRespository {
	return &postgresUserRespository{
		db: db,
	}
}

func (r *postgresUserRespository) Create(name, email string) (domain.User, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`

	var id string
	err := r.db.QueryRow(query, name, email).Scan(&id)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:    id,
		Name:  name,
		Email: email,
	}, nil
}

func (r *postgresUserRespository) List() ([]domain.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return []domain.User{}, err
	}
	defer rows.Close()
	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (r *postgresUserRespository) FindByID(id string) (domain.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	var user domain.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}
