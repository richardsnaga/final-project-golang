package repository

import (
	"database/sql"
	"final-project-golang/structs"
)

func GetAllChapter(db *sql.DB) (err error, results []structs.Chapter) {
	sql := `SELECT * FROM chapters`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var chapter = structs.Chapter{}

		err = rows.Scan(&chapter.Id, &chapter.ComicID, &chapter.ChapterNumber, &chapter.ImageUrl, &chapter.CreatedAt, &chapter.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, chapter)
	}

	return
}

func CreateChapter(db *sql.DB, chapter structs.Chapter) (err error) {
	sql := `INSERT INTO chapters (comic_id, chapter_number, image_url) VALUES ($1,$2,$3)`

	errs := db.QueryRow(sql, chapter.ComicID, chapter.ChapterNumber, chapter.ImageUrl)

	return errs.Err()
}

func UpdateChapter(db *sql.DB, chapter structs.Chapter) (err error) {
	sql := `UPDATE chapters SET comic_id = $1, chapter_number = $2, image_url = $3, updated_at = NOW()::timestamp WHERE id = $4`

	errs := db.QueryRow(sql, chapter.ComicID, chapter.ChapterNumber, chapter.ImageUrl, chapter.Id)

	return errs.Err()
}


func DeleteChapter(db *sql.DB, chapter structs.Chapter) (err error) {
	sql := `DELETE FROM chapters WHERE id = $1`

	errs := db.QueryRow(sql, chapter.Id)

	return errs.Err()
}