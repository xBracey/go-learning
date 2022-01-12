package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	routes "example.com/example/routes"
)

func Check(err error) {
	fmt.Println(err)

	if err != nil {
		log.Fatal(err)
	}
}

func AddTable(db *sql.DB) sql.Result {
	query := `
	CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username TEXT NOT NULL,
        password TEXT NOT NULL
    );
`

	result, _ := db.Exec(query)

	return result
}

func PostUser(db *sql.DB) sql.Result {
	query := `
	INSERT INTO users(username, password)
	VALUES('test123', 'test123');
`

	result, _ := db.Exec(query)

	return result
}

func main() {
	os.Setenv("MOOCDSN", "postgres://bucketeer:bucketeer_pass@db/bucketeer_db?sslmode=disable")

	db, err := sql.Open("postgres", os.Getenv("MOOCDSN"))
	Check(err)
	AddTable(db)

	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/users/post", func(w http.ResponseWriter, r *http.Request) {
		PostUser(db)
		fmt.Fprint(w, "Hello, you've added a user")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := routes.GetUsers(db)

		fmt.Fprint(w, users)
	})

	http.ListenAndServe(":8080", nil)
}
