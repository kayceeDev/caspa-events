package storage

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	"github.com/kayceeDev/caspa-events/ent-go/ent"
	"github.com/kayceeDev/caspa-events/ent-go/ent/migrate"
	ports "github.com/kayceeDev/caspa-events/internal/shared/ports/storage"

	"github.com/kayceeDev/caspa-events/pkg/config"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Database struct {
	client *ent.Client
}

var _ ports.IDataBase = (*Database)(nil)

func NewDatabase(cfg Config) (*Database, error) {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	drv, err := sql.Open(dialect.Postgres, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
	}

	db := drv.DB()
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(10000)
	db.SetConnMaxLifetime(30 * time.Second)
	db.SetConnMaxIdleTime(5 * time.Second)

	client := ent.NewClient(ent.Driver(drv))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = client.ExecContext(ctx, "CREATE EXTENSION IF NOT EXISTS vector")
	if err != nil {
		panic(err)
	}
	_, err = client.ExecContext(ctx, "CREATE EXTENSION IF NOT EXISTS pg_trgm")
	if err != nil {
		panic(err)
	}

	if err := client.Schema.Create(context.Background(), migrate.WithDropColumn(true), migrate.WithDropIndex(true)); // migrate.WithDropColumn(true), migrate.WithDropIndex(true)
	err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	return &Database{client: client}, nil
}

func (d *Database) Client() *ent.Client {
	return d.client
}

func (d *Database) Close() error {
	return d.client.Close()
}

type IGetEnv interface {
	GetEnv(key string, fallback string) string
}

func DefaultConfig() Config {
	env := config.EnvVars()

	port := env.DB_PORT
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("failed to convert port to integer: %v", err)
	}

	return Config{
		Host:     env.DB_HOST,
		Port:     portInt,
		User:     env.DB_USER,
		Password: env.DB_PASSWORD,
		DBName:   env.DB_NAME,
		SSLMode:  env.DB_SSLMODE,
	}
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			log.Println("panic occurred:", v)
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			log.Printf("rolling back transaction failed: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction failed: %w", err)
	}
	return nil
}

func (d *Database) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := d.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			log.Println("panic occurred:", v)
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			log.Printf("rolling back transaction failed: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction failed: %w", err)
	}
	return nil
}
