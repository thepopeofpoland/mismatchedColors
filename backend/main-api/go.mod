module mismatchedColors/api

go 1.22.0

require (
	github.com/gorilla/mux v1.8.1
	github.com/mattn/go-sqlite3 v1.14.22
)

require (
	mismatchedColors/database v0.0.0-00010101000000-000000000000
	mismatchedColors/models v0.0.0-00010101000000-000000000000
)

replace mismatchedColors/database => ../database

replace mismatchedColors/models => ../models
