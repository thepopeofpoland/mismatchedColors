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

func InsertNewRandomColor(db *sql.DB, randomColor models.RandomColors) models.RandomColors {
	colorSQL := "INSERT INTO RandomColors (ColorName, ColorHex) VALUES (?, ?)"
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


// --- Color Palettes ---
func GetAllColorPalettes(db *sql.DB) []models.ColorPalettes {
	var palettes []models.ColorPalettes

	row, err := db.Query("SELECT * FROM ColorPalettes")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var id int
		var name string
		var color1 int
		var color2 int
		var color3 int
		var color4 int
		var color5 int

		row.Scan(&id, &name, &color1, &color2, &color3, &color4, &color5)
		palette := models.ColorPalettes{
			PaletteID: id,
			PaletteName: name,
			Color1: color1,
			Color2: color2,
			Color3: color3,
			Color4: color4,
			Color5: color5,
		}
		palettes = append(palettes, palette)
	}
	return palettes
}


func InsertNewColorPalette(db *sql.DB, colorPalette models.RandomColors) models.RandomColors {
	paletteSQL := "INSERT INTO ColorPalettes (PaletteName, Color1, Color2, Color3, Color4, Color5) VALUES (?, ?, ?, ?, ?, ?)"
	statement, err := db.Prepare(paletteSQL)
	if err != nil {
		log.Fatal(err)
	}
	res, err := statement.Exec(colorPalette.PaletteName, colorPalette.Color1, colorPalette.Color2, colorPalette.Color3, colorPalette.Color4, colorPalette.Color5)
	if err != nil {
		log.Fatal(err)
	}
	lid, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	colorPalette.PaletteIDID = int(lid)
	return colorPalette
}
