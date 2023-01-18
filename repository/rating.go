package repository

import (
	"database/sql"
	"final-project-golang/structs"
)

func CreateRating(db *sql.DB, rating structs.Rating) (err error) {
	sql := `INSERT INTO rating (comic_id, user_id, rate) VALUES ($1,$2,$3)`

	errs := db.QueryRow(sql, rating.ComicId, rating.UserId, rating.Rate)

	return errs.Err()
}

func GetRatingByComicId(db *sql.DB, id int) (err error, results []structs.AvgRating) {
	sql := `SELECT comic_id, ROUND(AVG(rate),2) FROM rating where comic_id = $1 group by comic_id`

	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var rate = structs.AvgRating{}

		err = rows.Scan(&rate.ComicId, &rate.AvgRate)
		if err != nil {
			panic(err)
		}

		results = append(results, rate)
	}

	return
}
