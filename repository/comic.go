package repository

import (
	"database/sql"
	"final-project-golang/structs"
)

func GetAllComic(db *sql.DB) (err error, results []structs.Comic) {
	sql := `SELECT * FROM comic`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var comic = structs.Comic{}

		err = rows.Scan(&comic.Id, &comic.Title, &comic.Description, &comic.ImageURL, &comic.ReleaseYear, &comic.Genre, &comic.Type, &comic.Status, &comic.CreatedAt, &comic.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, comic)
	}

	return
}

func CreateComic(db *sql.DB, comic structs.Comic) (err error) {
	sql := `INSERT INTO comic (title, description, image_url, release_year, genre, type, status) VALUES ($1,$2,$3,$4,$5,$6,$7)`

	errs := db.QueryRow(sql, comic.Title, comic.Description, comic.ImageURL, comic.ReleaseYear, comic.Genre, comic.Type, comic.Status)

	return errs.Err()
}

func UpdateComic(db *sql.DB, comic structs.Comic) (err error) {
	sql := `UPDATE comic SET title = $1, description = $2, image_url = $3, release_year = $4, genre = $5, type = $6, status = $7, updated_at = NOW()::timestamp WHERE id = $8`

	errs := db.QueryRow(sql, comic.Title, comic.Description, comic.ImageURL, comic.ReleaseYear, comic.Genre, comic.Type, comic.Status, comic.Id)

	return errs.Err()
}

func DeleteComic(db *sql.DB, comic structs.Comic) (err error) {
	sql := `DELETE FROM comic WHERE id = $1`

	errs := db.QueryRow(sql, comic.Id)

	return errs.Err()
}

func FilterComic(db *sql.DB, genre string, tipe string, status bool) (err error, results []structs.Comic) {
	sql := `SELECT * FROM comic WHERE genre = $1 AND type = $2 AND status = $3`

	rows, err := db.Query(sql, genre, tipe, status)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var comic = structs.Comic{}

		err = rows.Scan(&comic.Id, &comic.Title, &comic.Description, &comic.ImageURL, &comic.ReleaseYear, &comic.Genre, &comic.Type, &comic.Status, &comic.CreatedAt, &comic.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, comic)
	}

	return
}