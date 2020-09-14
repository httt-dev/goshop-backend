package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"timewise/model"
)

//SQL object keep connect to database + provide query function
type SQL struct {
	Db *sqlx.DB
}

// connect to postgres db
func (s *SQL) Connect(cfg model.Config) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.DbHost,
		cfg.Database.DbPort,
		cfg.Database.DbUserName,
		cfg.Database.DbPassword,
		cfg.Database.DbName)

	//fmt.Println(dataSource)

	s.Db = sqlx.MustConnect("postgres", dataSource)
	if err := s.Db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect db ok")
}

//close connection
func (s *SQL) Close() {
	s.Db.Close()
}
