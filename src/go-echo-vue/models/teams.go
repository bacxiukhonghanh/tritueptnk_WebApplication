package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Team is a struct containing Team data
type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Token string `json:"token"`
	Link string `json:"link"`
}

// TeamCollection is collection of Teams
type TeamCollection struct {
	Teams []Team `json:"items"`
}

// GetTeams from the DB
func GetTeams(db *sql.DB) TeamCollection {
	sql := "SELECT * FROM teams"
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := TeamCollection{}
	for rows.Next() {
		team := Team{}
		err2 := rows.Scan(&team.ID, &team.Name, &team.Token, &team.Link)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Teams = append(result.Teams, team)
	}
	return result
}

// PutTeam into DB
func PutTeam(db *sql.DB, name string, token string, link string) (int64, error) {
	sql := "INSERT INTO teams(name, token, link) VALUES(?, ?, ?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(name, token, link)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeleteTeam from DB
func DeleteTeam(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM teams WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
