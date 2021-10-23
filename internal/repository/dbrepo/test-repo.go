package dbrepo

import (
	"errors"

	"github.com/DungBuiTien1999/udemy-api/internal/models"
)

// GetTasksByUserID returns a slice of tasks by user id
func (m *testDBRepo) GetTasksByUserID(userID int) ([]models.Task, error) {
	var tasks []models.Task

	if userID == 2 {
		return tasks, errors.New("some errors")
	}

	return tasks, nil
}

// InsertTask inserts a task into database
func (m *testDBRepo) InsertTask(task models.Task) error {
	if task.UserID == 1 {
		return nil
	}
	return errors.New("some errors")
}

// UpdateTaskByID updates a task by id
func (m *testDBRepo) UpdateTaskByID(id int, task models.Task) (models.Task, error) {
	var newTask models.Task
	if id == 1 {
		return newTask, nil
	}

	return newTask, errors.New("some errors")
}

// GetUserByUsername returns an user by username
func (m *testDBRepo) GetUserByUsername(username string) (models.User, error) {
	user := models.User{
		ID:       2,
		Username: "samurai",
		Password: " $2a$10$n.zkqF4h3942tS2nBrisKuwSn/s8J5Y9pZMCsMWkKMh.Ewwsl7xei",
	}
	user2 := models.User{
		ID:       1,
		Username: "tiendung",
		Password: "$2a$10$ReN0HKWS9fghNvpbpogGte09dSilOs2bFdniqrgvIS2Z02y.irkty",
	}
	if username == user.Username {
		return user, nil
	}
	if username == user2.Username {
		return user2, nil
	}

	return user, errors.New("some errors")
}

// InsertUser inserts an user into database
func (m *testDBRepo) InsertUser(user models.User) error {
	if user.Username == "dungbui" {
		return nil
	}
	return errors.New("some errors")
}

// UpdateRefreshToken updates refresh token by user id
func (m *testDBRepo) UpdateRefreshToken(id int, rfToken string) error {
	if id == 1 {
		return nil
	}
	return errors.New("some errors")
}

// IsValidRefreshToken return if refresh token is valid with user id
func (m *testDBRepo) IsValidRefreshToken(id int, rfToken string) bool {

	return true
}
