package routes

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type user struct {
	id       int
	username string
	password string
}

func GetUsers(db *sql.DB) []user {
	users := []user{}

	query := `
	SELECT	* FROM users;
`

	rows, _ := db.Query(query)

	defer rows.Close()

	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	return users
}
