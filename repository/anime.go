package repository

import (
	"database/sql"
	"final-project-go-sanber/structs"
)

func GetAllAnime(db *sql.DB) (result []structs.Anime, err error) {
	sql := "SELECT * FROM anime"
	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var anime structs.Anime

		err = rows.Scan(&anime.ID, &anime.Title, &anime.Description, &anime.Genre, &anime.ReleaseYear)
		if err != nil {
			return
		}

		result = append(result, anime)
	}
	return
}

func InsertAnime(db *sql.DB, anime structs.Anime) (err error) {
	sql := "INSERT INTO anime(id, title, description_anime, genre, release_year) VALUES ($1, $2, $3, $4, $5)"

	errs := db.QueryRow(sql, anime.ID, anime.Title, anime.Description, anime.Genre, anime.ReleaseYear)

	return errs.Err()
}

func UpdateAnime(db *sql.DB, anime structs.Anime) (err error) {
	sql := "UPDATE anime SET title = $1, description_anime = $2, genre = $3, release_year = $4 WHERE id = $5"

	errs := db.QueryRow(sql, anime.Title, anime.Description, anime.Genre, anime.ReleaseYear, anime.ID)

	return errs.Err()
}

func DeleteAnime(db *sql.DB, anime structs.Anime) (err error) {
	sql := "DELETE FROM anime WHERE id = $1"

	errs := db.QueryRow(sql, anime.ID)
	return errs.Err()
}
