package database

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type DBManager struct {
	DBPool *pgxpool.Pool
}

func (m *DBManager) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := m.DBPool.Acquire(context.Background())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer conn.Release()

		// Set the connection as a value in the context
		c.Set("DBManager", conn)

		c.Next()
	}
}
