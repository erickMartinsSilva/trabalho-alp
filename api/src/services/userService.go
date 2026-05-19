package services

import (
	"ALP/src/db"
	"ALP/src/models"
)

func CreateUser(user models.User) error {
	query := "INSERT INTO users VALUES (?)"
	_, err := db.Instance.Exec(query, user.Name)
	return err
}

func GetUserById(id int) (models.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.Instance.QueryRow(query, id)
	
	var user models.User

	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUsers() ([]models.User, error) {
	query := "SELECT * FROM users"
	rows, error := db.Instance.Query(query)

	if error != nil {
		return nil, error
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUser(id int, data models.User) error {
	query := "UPDATE users SET name = ? WHERE id = ?"
	_, err := db.Instance.Exec(query, data.Name, id)
	return err
}

func DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Instance.Exec(query, id)
	return err
}