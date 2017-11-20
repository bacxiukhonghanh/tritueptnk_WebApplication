package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"go-echo-vue/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// GetTeams endpoint
func GetTeams(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTeams(db))
	}
}

// PutTeam endpoint
func PutTeam(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new team
		var team models.Team
		// Map imcoming JSON body to the new Team
		c.Bind(&team)
		// Add a team using our new model
		id, err := models.PutTeam(db, team.Name, team.Token, team.Link)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// Handle any errors
		} else {
			return err
		}
	}
}

// DeleteTeam endpoint
func DeleteTeam(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a team
		_, err := models.DeleteTeam(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
			// Handle errors
		} else {
			return err
		}
	}
}
