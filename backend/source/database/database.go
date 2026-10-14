package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/F4nk1/agentrix/backend/source/config"
	"github.com/F4nk1/agentrix/backend/source/tracer"
)

type Database struct {
	Db *sql.DB
}

func NewDataBase(ctx context.Context, config *config.Config) (*Database, error) {
	tracer.Debug(ctx, "Initializing database connection")

	connectionString := config.GetStringDBConnection()

	db, err := sql.Open(config.Database.Driver, connectionString)
	if err != nil {
		tracer.Errorf(ctx, "Failed to open database connectoin: %s", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		tracer.Errorf(ctx, "Failed to ping database: %s", err)
		return nil, fmt.Errorf("Cannot connect to database: %v", err)
	}

	tracer.Debug(ctx, "Connected to database successfully")

	database := &Database{
		Db: db,
	}

	return database, nil
}

func (c *Database) Close() error {
	return c.Db.Close()
}

