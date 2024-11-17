package storage

import (
	"database/sql"
)

type PostgreSQL struct {
	db *sql.DB
}

// NewPostgresql инициализирует объект PostgreSQL с заданной конфигурацией подключения.
// Возвращает указатель на PostgreSQL.
func NewPostgresql(cnf string) *PostgreSQL {
	var pstg = PostgreSQL{}
	//pstg.config = cnf
	return &pstg
}

// Connect устанавливает подключение к базе данных.
// Если подключение успешно, возвращает nil, в противном случае возвращает ошибку.
func (pstg *PostgreSQL) Connect() error {
	/*	var err error
		pstg.db, err = sql.Open("pgx", pstg.config)
		if err != nil {
			log.Println("error connecting to Postgresql ", err)
			return err
		}

		err = pstg.PingContext()
		if err != nil {
			log.Println("error pinging Postgresql ", err)
			return err
		}*/

	return nil
}
