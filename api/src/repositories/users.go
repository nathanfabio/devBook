package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// NewUsersRepositories creates a user repository
func NewUsersRepositories(db *sql.DB) *users {
	return &users{db}
}

// Create inserts a user into the database
func (u users) Create(user models.User) (uint64, error) {
	statment, err := u.db.Prepare("insert into users (name, nick, email, password) values(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastEnteredID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastEnteredID), nil
}
