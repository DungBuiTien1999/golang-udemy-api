package dbrepo

import (
	"context"
	"time"

	"github.com/DungBuiTien1999/udemy-api/internal/models"
)

// GetTasksByUserID returns a slice of tasks by user id
func (m *mysqlDBRepo) GetTasksByUserID(userID int) ([]models.Task, error) {
	// request last longer 3 second so discard write record into db
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, complete, user_id, created_at, updated_at from tasks where user_id = ?`

	var tasks []models.Task

	rows, err := m.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Complete,
			&task.UserID,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}

// InsertTask inserts a task into database
func (m *mysqlDBRepo) InsertTask(task models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into tasks 
	(title, complete, user_id) values 
	(?, ?, ?)
	`

	_, err := m.DB.ExecContext(ctx, stmt, task.Title, task.Complete, task.UserID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTaskByID updates a task by id
func (m *mysqlDBRepo) UpdateTaskByID(id int, task models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update tasks set title = ?, complete = ?, updated_at = ? where id = ?`

	var newTask models.Task

	_, err := m.DB.ExecContext(ctx, query, task.Title, task.Complete, time.Now(), id)
	if err != nil {
		return newTask, err
	}
	newTask.ID = id
	newTask.Title = task.Title
	newTask.Complete = task.Complete
	newTask.UpdatedAt = time.Now()

	return newTask, nil
}

// GetUserByUsername returns an user by username
func (m *mysqlDBRepo) GetUserByUsername(username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, username, password, created_at, updated_at from users where username = ?`

	var user models.User

	err := m.DB.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

// InsertUser inserts an user into database
func (m *mysqlDBRepo) InsertUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into users 
	(username, password) values 
	(?, ?)
	`

	_, err := m.DB.ExecContext(ctx, stmt, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRefreshToken updates refresh token by user id
func (m *mysqlDBRepo) UpdateRefreshToken(id int, rfToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update users set rfToken = ? where id = ?`

	_, err := m.DB.ExecContext(ctx, query, rfToken, id)
	if err != nil {
		return err
	}
	return nil
}

// IsValidRefreshToken return if refresh token is valid with user id
func (m *mysqlDBRepo) IsValidRefreshToken(id int, rfToken string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select rfToken from users where id = ?`

	var user models.User

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&user.RfToken)
	if err != nil {
		return false
	}
	return rfToken == user.RfToken
}
