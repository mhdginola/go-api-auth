package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func InitDB() {
	start := time.Now()

	// Database connection configuration
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		log.Fatal("Error parsing connection config: ", err)
	}

	dbPort, errPort := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32)

	if err != nil {
		fmt.Println("Invalid input:", errPort)
	}

	config.ConnConfig.Host = os.Getenv("DB_HOST")
	config.ConnConfig.Port = uint16(dbPort)
	config.ConnConfig.User = os.Getenv("DB_USERNAME")
	config.ConnConfig.Password = os.Getenv("DB_PASSWORD")
	config.ConnConfig.Database = os.Getenv("DB_NAME")

	// Set maximum number of connections in the pool
	config.MinConns = 5
	config.MaxConns = 10
	config.MaxConnIdleTime = 5 * time.Minute
	config.MaxConnLifetime = 10 * time.Minute

	// Create a connection pool
	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err)
	}

	db = dbpool
	fmt.Printf("initial DB took %v\n", time.Since(start))
}

func Conn() *pgxpool.Pool {
	return db
}

var GetDBConn = func() PgxConnIface {
	return Conn()
}

type PgxConnIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	Close()
}
