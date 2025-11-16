package auth

import (
	"context"
	"database/sql"
	"ginauth/src/config"
	"ginauth/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler authenticates and returns a JWT
func LoginHandler(c *gin.Context) {
	conn := config.GetDBConn()
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	var id int
	var hashed string
	row := conn.QueryRow(ctx, "SELECT id, password FROM users WHERE username=$1", req.Username)
	switch err := row.Scan(&id, &hashed); err {
	case nil:
		// ok
	case sql.ErrNoRows:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	if hashed != utils.Sha256Hash(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// create JWT
	tokenString, err := utils.CreateToken(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
