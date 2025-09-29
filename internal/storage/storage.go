package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
	config *Config
}

type Config struct {
	Host string
	Port string
	Password string
	User string
	Database string
	SSLMode string
}


func NewStorage(cfg *Config) *Storage {
	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable"
	}

	return &Storage{
		config: cfg,
	}
}

// TODO: добавить информацию о подключении в logger
func (s *Storage) Connect() error {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		s.config.User,
		s.config.Password,
		s.config.Host,
		s.config.Port,
		s.config.Database,
		s.config.SSLMode,
	)

	config, err := pgxpool.ParseConfig(connectionString)
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	if err != nil {
		panic(fmt.Errorf("Не удалось подключиться к базе данных: %s", err))
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		panic(fmt.Errorf("Не удалось подключиться к базе данных: %s", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	s.pool = pool
	
	return nil

}

func (s *Storage) Close() {
	if s.pool != nil {
		s.pool.Close()
	}
}

func (s *Storage) Pool() *pgxpool.Pool {
	return s.pool
}

func (s *Storage) Start() error {
	if err := s.Connect(); err != nil {
		return err
	}

	

	return nil
}
