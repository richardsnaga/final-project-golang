package repository

import (
	"database/sql"
	"final-project-golang/structs"
)

func GetAllComment(db *sql.DB) (err error, results []structs.Comment) {
	sql := `SELECT * FROM comments`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var comment = structs.Comment{}

		err = rows.Scan(&comment.Id, &comment.ChapterId, &comment.UserId, &comment.ReferenceId, &comment.Comment, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, comment)
	}

	return
}

func CreateComment(db *sql.DB, comment structs.Comment) (err error) {
	sql := `INSERT INTO comments (chapter_id, user_id, reference_id, comment) VALUES ($1,$2,$3,$4)`

	errs := db.QueryRow(sql, comment.ChapterId, comment.UserId, comment.ReferenceId, comment.Comment)

	return errs.Err()
}

func UpdateComment(db *sql.DB, comment structs.Comment) (err error) {
	sql := `UPDATE comments SET chapter_id = $1, user_id = $2, reference_id = $3, comment = $4, updated_at = NOW()::timestamp WHERE id = $5`

	errs := db.QueryRow(sql, comment.ChapterId, comment.UserId, comment.ReferenceId, comment.Comment, comment.Id)

	return errs.Err()
}

func DeleteComment(db *sql.DB, comment structs.Comment) (err error) {
	sql := `DELETE FROM comments WHERE id = $1`

	errs := db.QueryRow(sql, comment.Id)

	return errs.Err()
}

func GetCommentByChapterId(db *sql.DB, id int) (err error, results []structs.Comment) {
	sql := `SELECT * FROM comments WHERE chapter_id = $1`

	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var comment = structs.Comment{}

		err = rows.Scan(&comment.Id, &comment.ChapterId, &comment.UserId, &comment.ReferenceId, &comment.Comment, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, comment)
	}

	return
}
