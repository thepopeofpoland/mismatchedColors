package colorDB

import (
	"database/sql"
	"log"
	"mismatchedColors/models"
)

func GetAllRandomColors(db *sql.DB) []models.RandomColors {
	var colors []models.RandomColors

	row, err := db.Query("SELECT * FROM RandomColors")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var id int
		var name string
		var hex string

		row.Scan(&id, &name, &hex)
		randomColors := models.RandomColors{
			ColorID:   id,
			ColorName: name,
			ColorHex:  hex,
		}
		colors = append(colors, randomColors)
	}
	return colors
}

func insertNewRandomColor(db *sql.DB, randomColor models.RandomColors) models.RandomColors {
	colorSQL := "INSERT INTO RandomColors (name, hex) VALUES (?, ?)"
	statement, err := db.Prepare(colorSQL)
	if err != nil {
		log.Fatal(err)
	}
	res, err := statement.Exec(randomColor.ColorName, randomColor.ColorHex)
	if err != nil {
		log.Fatal(err)
	}
	lid, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	randomColor.ColorID = int(lid)
	return randomColor
}