package main

import (
	"database/sql"

	"go-echo-vue/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := initDB("storage.db")
	migrate(db)

	e := echo.New()

	e.File("/TeamManager", "public/TeamManager.html")
	e.GET("/teams", handlers.GetTeams(db))
	e.PUT("/teams", handlers.PutTeam(db))
	e.DELETE("/teams/:id", handlers.DeleteTeam(db))

	e.Run(standard.New(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS teams(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		token VARCHAR NOT NULL,
		link VARCHAR NOT NULL
	);
	CREATE TABLE IF NOT EXISTS questions(
		quesid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		token VARCHAR NOT NULL,
		ans VARCHAR NOT NULL
	);
	CREATE TABLE IF NOT EXISTS submissions(
		token VARCHAR NOT NULL,
		quesid INTEGER NOT NULL,		
		submission VARCHAR NOT NULL
	);
	`

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
