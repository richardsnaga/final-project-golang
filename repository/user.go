package repository

import (
	"database/sql"
	"final-project-golang/structs"
)

func Register(db *sql.DB, user structs.User) (err error) {
	sql := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`

	errs := db.QueryRow(sql, user.Name, user.Email, user.Password)

	return errs.Err()
}

// func Login(db *sql.DB, user structs.User) (count int) {
// 	var count int
// 	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email=$1 AND password=$2", json.Email, json.Password).Scan(&count)
// 	return count
// }

func UpdateUser(db *sql.DB, user structs.User) (err error) {
	sql := `UPDATE users SET name = $1, email = $2, password = $3, updated_at = NOW()::timestamp WHERE id = $4`

	errs := db.QueryRow(sql, user.Name, user.Email, user.Password, user.Id)

	return errs.Err()
}

func DeleteUser(db *sql.DB, user structs.User) (err error) {
	sql := `DELETE FROM users WHERE id = $1`

	errs := db.QueryRow(sql, user.Id)

	return errs.Err()
}
