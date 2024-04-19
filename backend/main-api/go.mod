module mismatchedColors/api

go 1.22.0

require (
	github.com/gorilla/mux v1.8.1
	github.com/mattn/go-sqlite3 v1.14.22
	mismatchedColors/database v0.0.0-00010101000000-000000000000
)

require mismatchedColors/models v0.0.0-00010101000000-000000000000 // indirect

replace mismatchedColors/database => ../database

replace mismatchedColors/models => ../models
