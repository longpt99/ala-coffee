package database

import (
	"context"
	"ecommerce/configs"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	DB  *pgxpool.Pool
	Ctx context.Context
}

// var (
// 	pgInstance *postgres
// 	pgOnce     sync.Once
// )

// func NewPG(ctx context.Context, connString string) (*postgres, error) {
// 	pgOnce.Do(func() {
// 		db, err := pgxpool.New(ctx, connString)
// 		if err != nil {
// 			// return nil, fmt.Errorf("unable to create connection pool: %w", err)
// 			return
// 		}

// 		pgInstance = &postgres{db}
// 	})

// 	return pgInstance, nil
// }

// func (pg *postgres) Ping(ctx context.Context) error {
// 	return pg.DB.Ping(ctx)
// }

func (pg *PostgresConfig) Close() {
	pg.DB.Close()
}

func InitDatabase(ctx context.Context) *PostgresConfig {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		configs.Env.Postgres.Username,
		configs.Env.Postgres.Password,
		configs.Env.Postgres.Host,
		configs.Env.Postgres.Port,
		configs.Env.Postgres.Database,
	)

	db, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	if err = db.Ping(ctx); err != nil {
		log.Fatalf("Ping database error: %v\n", err)
	}

	err = runSQLFiles(db, "/configs/database/migrates")

	if err != nil {
		log.Fatal(err)
	}

	return &PostgresConfig{db, ctx}
}

func runSQLFiles(conn *pgxpool.Pool, path string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	dirPath := filepath.Join(wd, path)
	files, err := os.ReadDir(dirPath)

	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			filePath := filepath.Join(dirPath, file.Name())
			if err := runSQLFile(conn, filePath); err != nil {
				return err
			}
		}
	}

	return nil
}

func runSQLFile(conn *pgxpool.Pool, filePath string) error {
	sqlFile, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading SQL file:", err)
		return err
	}

	_, err = conn.Exec(context.Background(), string(sqlFile))
	if err != nil {
		fmt.Println("Error executing SQL:", err)
		return err
	}

	return nil
}
