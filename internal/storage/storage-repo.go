package storage

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"strings"
	"user-service/internal/storage/migrations"
)

type DBConnection struct {
	Connection
	connection *gorm.DB
}

type DBSession struct {
	Session
	session *gorm.DB
	log     *slog.Logger
	v       *validator.Validate // validator for incoming data(internal layer)
}

func (c *DBConnection) NewSession(log *slog.Logger) *DBSession {
	return &DBSession{
		session: c.connection.Session(&gorm.Session{}),
		log:     log,
	}
}

func MustLoadDB(host string, port string, user string, password string, dbname string) *DBConnection {
	var dbPath strings.Builder

	dbPath.WriteString("host=")
	dbPath.WriteString(host)
	dbPath.WriteString(" port=")
	dbPath.WriteString(port)
	dbPath.WriteString(" user=")
	dbPath.WriteString(user)
	dbPath.WriteString(" password=")
	dbPath.WriteString(password)
	dbPath.WriteString(" dbname=")
	dbPath.WriteString(dbname)
	dbPath.WriteString(" sslmode=disable")

	db, err := gorm.Open(postgres.Open(dbPath.String()), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	migrations.MustMigrate(db)

	return &DBConnection{connection: db}
}
