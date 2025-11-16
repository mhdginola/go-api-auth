package auth

import (
	"context"
	"database/sql"
	"ginauth/src/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfileHandler(c *gin.Context) {
	conn := config.GetDBConn()
	sub, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no user in context"})
		return
	}
	userID := sub.(string)

	ctx := context.Background()
	var username string
	row := conn.QueryRow(ctx, "SELECT username FROM users WHERE id=$1", userID)
	switch err := row.Scan(&username); err {
	case nil:
		c.JSON(http.StatusOK, gin.H{"id": userID, "username": username})
	case sql.ErrNoRows:
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
	}
}
