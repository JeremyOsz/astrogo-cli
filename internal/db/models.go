package db

import (
	"time"
)

// User represents a user in the database
type User struct {
	ID        int64
	Username  string
	Email     string
	CreatedAt time.Time
}

// CreateUser inserts a new user into the database
func CreateUser(username, email string) (*User, error) {
	result, err := DB.Exec(`
		INSERT INTO users (username, email)
		VALUES (?, ?)
	`, username, email)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetUserByID(id)
}

// GetUserByID retrieves a user by their ID
func GetUserByID(id int64) (*User, error) {
	user := &User{}
	err := DB.QueryRow(`
		SELECT id, username, email, created_at
		FROM users
		WHERE id = ?
	`, id).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByUsername retrieves a user by their username
func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	err := DB.QueryRow(`
		SELECT id, username, email, created_at
		FROM users
		WHERE username = ?
	`, username).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]*User, error) {
	rows, err := DB.Query(`
		SELECT id, username, email, created_at
		FROM users
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// UpdateUser updates an existing user's information
func UpdateUser(id int64, username, email string) (*User, error) {
	_, err := DB.Exec(`
		UPDATE users
		SET username = ?, email = ?
		WHERE id = ?
	`, username, email, id)
	if err != nil {
		return nil, err
	}

	return GetUserByID(id)
}

// DeleteUser removes a user from the database
func DeleteUser(id int64) error {
	_, err := DB.Exec(`
		DELETE FROM users
		WHERE id = ?
	`, id)
	return err
}
