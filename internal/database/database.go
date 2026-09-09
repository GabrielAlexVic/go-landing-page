package database

import (
	"database/sql"
	"fmt"
	"go-landing-page/internal/database/queries"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/vinovest/sqlx"
)

type Service interface {
	Health() map[string]string
	NewTx() (*sql.Tx, error)
	GetDB() *sqlx.DB
	Close() error
}

type service struct {
	db *sqlx.DB
}

func NewConnectionDatabase() *service {
	connectionString := getConnectionString()

	db := sqlx.MustConnect("postgres", connectionString)
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	log.Println("Conectado ao banco de dados com sucesso!")

	err := createAllTables(db)
	if err != nil {
		panic(err.Error())
	}

	log.Println("Tabelas sincronizadas com sucesso!")

	return &service{
		db: db,
	}
}

func getConnectionString() string {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)
}

func createAllTables(db *sqlx.DB) error {
	_, err := db.Exec(queries.CreateUserTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(queries.CreatePageViewsTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(queries.CreateLeadsTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(queries.CreateEventsTable)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Close() error {
	return s.db.Close()
}

func (s *service) Health() map[string]string {
	err := s.db.Ping()
	if err != nil {
		return map[string]string{
			"database": "unhealthy",
		}
	}
	return map[string]string{
		"database": "healthy",
	}
}

func (s *service) NewTx() (*sql.Tx, error) {
	return s.db.Begin()
}

func (s *service) GetDB() *sqlx.DB {
	return s.db
}
