package auth

import (
	"context"
	"ginauth/src/config"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	conn := config.GetDBConn()

	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		RoleID   int    `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	// Check if user exists
	var exists bool
	err := conn.QueryRow(ctx,
		"SELECT EXISTS (SELECT 1 FROM users WHERE username=$1)",
		req.Username,
	).Scan(&exists)
	if err != nil {
		c.JSON(500, gin.H{"error": "db error"})
		return
	}

	if exists {
		c.JSON(409, gin.H{"error": "username already exists"})
		return
	}

	// Insert user
	var id int
	err = conn.QueryRow(ctx,
		"INSERT INTO users (username, password, role_id) VALUES ($1, encode(digest($2, 'sha256'), 'hex'), $3) RETURNING id",
		req.Username, req.Password, req.RoleID,
	).Scan(&id)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(201, gin.H{
		"id":       id,
		"username": req.Username,
		"role_id":  req.RoleID,
	})
}
