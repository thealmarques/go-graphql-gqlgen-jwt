package services

import (
	"go-graphql-jwt/graph/injection"
	"go-graphql-jwt/graph/model"
)

// CreateUser - saves user in the DB
func CreateUser(user *model.User) error {
	db := injection.DB

	result, err := db.Exec("INSERT INTO `users` (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return err
	}

	user.ID = lastID
	return nil
}

// FindUsers - get all users from DB
func FindUsers() ([]*model.User, error) {
	db := injection.DB
	var users []*model.User

	rows, err := db.Query("SELECT id, name, email, password, created_at, updated_at FROM `users`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// FindUserByEmail - get user given an email
func FindUserByEmail(email string) (*model.User, error) {
	db := injection.DB
	row := db.QueryRow("SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?", email)
	var user model.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
