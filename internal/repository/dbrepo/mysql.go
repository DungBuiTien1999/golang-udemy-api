package dbrepo

import (
	"context"
	"time"

	"github.com/DungBuiTien1999/udemy-api/internal/models"
)

func (m *mysqlDBRepo) AllUsers() ([]models.User, error) {
	// request last longer 3 second so discard write record into db
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, first_name, last_name, created_at, updated_at from users
	`
	var users []models.User

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}
