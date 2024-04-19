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
