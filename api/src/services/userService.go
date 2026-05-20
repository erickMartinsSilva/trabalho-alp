package services

import (
	"ALP/src/db"
	"ALP/src/models"
)

func CreateUser(data models.User) (models.User, error) {
	query := "INSERT INTO users (name) VALUES (?) RETURNING id, name"
	res := db.Instance.QueryRow(query, data.Name)
	if res.Err() != nil {
		return models.User{}, res.Err()
	}

	var user = models.User{}
	res.Scan(&user.ID, &user.Name)
	return user, nil
}

func GetUserById(id string) (models.User, error) {
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

func UpdateUser(id string, data models.User) (models.User, error) {
	query := "UPDATE users SET name = ? WHERE id = ? RETURNING id, name"
	res := db.Instance.QueryRow(query, data.Name, id)

	var user models.User
	if err := res.Scan(&user.ID, &user.Name); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func DeleteUser(id string) (models.User, error) {
	query := "DELETE FROM users WHERE id = ? RETURNING id, name"
	res := db.Instance.QueryRow(query, id)

	var user models.User
	if err := res.Scan(&user.ID, &user.Name); err != nil {
		return models.User{}, err
	}
	return user, nil
}