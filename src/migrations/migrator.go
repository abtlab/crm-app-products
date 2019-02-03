package migrator

import (
	"database/sql"
	_ "github.com/lib/pq"
	"io/ioutil"
)

// Run :
func Run(connectionString string) {
	db, error := sql.Open("postgres", connectionString)
	if error != nil {
		panic(error)
	}

	InitDb(db)
	Migrate(db)

	defer db.Close()
}

// InitDb :
func InitDb(db *sql.DB) {
	const directory string = "migrations/init/"

	query := ReadFileAsString(directory + "create_db.sql")

	ExecuteQuery(db, query)
}

// Migrate :
func Migrate(db *sql.DB) {
	const directory string = "migrations/versions/"

	files, error := ioutil.ReadDir(directory)
	if error != nil {
		panic(error)
	}

	for _, file := range files {
		if !file.IsDir() {
			query := ReadFileAsString(directory + file.Name())

			ExecuteQuery(db, query)
		}
	}
}

// ExecuteQuery :
func ExecuteQuery(db *sql.DB, query string) {
	_, error := db.Exec(query)
	if error != nil {
		panic(error)
	}
}

// ReadFileAsString :
func ReadFileAsString(path string) string {
	bytes, error := ioutil.ReadFile(path)
	if error != nil {
		panic(error)
	}

	return string(bytes)
}
