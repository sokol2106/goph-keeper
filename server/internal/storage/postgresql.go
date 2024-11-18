package storage

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"server/internal/config"
	"server/internal/model"
)

type PostgreSQL struct {
	db     *sql.DB
	config config.Config
}

// NewPostgresql инициализирует объект PostgreSQL с заданной конфигурацией подключения.
// Возвращает указатель на PostgreSQL.
func NewPostgresql(config config.Config) *PostgreSQL {
	return &PostgreSQL{
		config: config,
	}
}

// Connect устанавливает подключение к базе данных.
// Если подключение успешно, возвращает nil, в противном случае возвращает ошибку.
func (pstg *PostgreSQL) Connect() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pstg.config.Postgres.Host,
		pstg.config.Postgres.Port,
		pstg.config.Postgres.User,
		pstg.config.Postgres.Password,
		pstg.config.Postgres.Database,
	)

	var err error
	pstg.db, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Println("Ошибка при подключении к PostgreSQL:", err)
		return err
	}

	err = pstg.db.Ping()
	if err != nil {
		log.Println("Ошибка при проверке подключения к PostgreSQL:", err)
		return err
	}

	log.Println("Успешное подключение к PostgreSQL")
	return nil
}

func (pstg *PostgreSQL) Close() error {
	return pstg.db.Close()
}

func (pstg *PostgreSQL) SelectUser(user model.User) (model.UserResponse, error) {
	query := `SELECT private_user_key FROM private_user WHERE login = $1 and password_hash = $2 `

	var PrivateUserKey uuid.UUID
	err := pstg.db.QueryRow(query, user.Login, user.PasswordHash).Scan(&PrivateUserKey)
	if err != nil {
		return model.UserResponse{}, err
	}

	return model.UserResponse{PrivateUserKey: PrivateUserKey}, nil
}

func (pstg *PostgreSQL) InsertUser(user model.User) (model.UserResponse, error) {
	query := `INSERT INTO private_user (login, password_hash, encryption_key) VALUES ($1, $2, $3) RETURNING private_user_key`

	var PrivateUserKey uuid.UUID
	err := pstg.db.QueryRow(query, user.Login, user.PasswordHash, user.EncryptionKey).Scan(&PrivateUserKey)
	if err != nil {
		return model.UserResponse{}, err
	}

	return model.UserResponse{PrivateUserKey: PrivateUserKey}, nil
}

func (pstg *PostgreSQL) InsertDataText(data model.DataText) (model.DataTextResponse, error) {
	query := `INSERT INTO data_text (private_user_key, data)
		VALUES ($1, $2) RETURNING data_text_key`

	var insertedUUID uuid.UUID
	err := pstg.db.QueryRow(query, data.PrivateUserKey, data.Data).Scan(&insertedUUID)
	if err != nil {
		return model.DataTextResponse{}, err
	}

	return model.DataTextResponse{DataTextKey: insertedUUID}, nil
}

func (pstg *PostgreSQL) SelectDataText(data model.DataText) (model.DataTextResponse, error) {
	query := `SELECT data_text_key, data
              FROM data_text
              WHERE data_text_key = $1 AND private_user_key = $2`

	var dataText model.DataTextResponse
	err := pstg.db.QueryRow(query, data.DataTextKey, data.PrivateUserKey).Scan(
		&dataText.DataTextKey,
		&dataText.Data,
	)

	if err != nil {
		return model.DataTextResponse{}, err
	}

	return dataText, nil
}

func (pstg *PostgreSQL) DeleteDataText(data model.DataText) error {
	query := `DELETE FROM data_text
              WHERE data_text_key = $1 AND private_user_key = $2`

	_, err := pstg.db.Exec(query, data.DataTextKey, data.PrivateUserKey)
	if err != nil {
		return err
	}

	return nil
}

func (pstg *PostgreSQL) InsertDataBinary(data model.DataBinary) (model.DataBinaryResponse, error) {
	query := `INSERT INTO data_binary (private_user_key, filename, data)
		VALUES ($1, $2, $3) RETURNING data_binary_key`

	var insertedUUID uuid.UUID
	binaryData := []byte(data.Data)
	err := pstg.db.QueryRow(query, data.PrivateUserKey, data.FileName, binaryData).Scan(&insertedUUID)
	if err != nil {
		return model.DataBinaryResponse{}, err
	}

	return model.DataBinaryResponse{DataBinaryKey: insertedUUID}, nil
}

func (pstg *PostgreSQL) SelectDataBinary(data model.DataBinary) (model.DataBinaryResponse, error) {
	query := `SELECT data_binary_key, filename, data
              FROM data_binary
              WHERE data_binary_key = $1 AND private_user_key = $2`

	var dataBinary model.DataBinaryResponse
	err := pstg.db.QueryRow(query, data.DataBinaryKey, data.PrivateUserKey).Scan(
		&dataBinary.DataBinaryKey,
		&dataBinary.FileName,
		&dataBinary.Data,
	)

	if err != nil {
		return model.DataBinaryResponse{}, err
	}

	return dataBinary, nil
}

func (pstg *PostgreSQL) DeleteDataBinary(data model.DataBinary) error {
	query := `DELETE FROM data_binary
              WHERE data_binary_key = $1 AND private_user_key = $2`

	_, err := pstg.db.Exec(query, data.DataBinaryKey, data.PrivateUserKey)
	if err != nil {
		return err
	}

	return nil
}

func (pstg *PostgreSQL) InsertDataCard(data model.DataCreditCard) (model.DataCreditCardResponse, error) {
	query := `INSERT INTO public.data_credit_cards (
                                      card_number, 
                                      cardholder_name, 
                                      expiration_date, 
                                      cvv_hash, 
                                      private_user_key) 
		VALUES ($1, $2, $3, $4, $5) RETURNING data_credit_card_key`

	var insertedUUID uuid.UUID

	err := pstg.db.QueryRow(
		query,
		data.CardNumber,
		data.CardholderName,
		data.ExpirationDate,
		data.CVVHash,
		data.PrivateUserKey,
	).Scan(&insertedUUID)
	if err != nil {
		return model.DataCreditCardResponse{}, err
	}

	return model.DataCreditCardResponse{
		DataCreditCardKey: insertedUUID,
	}, nil
}

func (pstg *PostgreSQL) SelectDataCard(data model.DataCreditCard) (model.DataCreditCardResponse, error) {
	query := `SELECT 
    	data_credit_card_key, 
       	card_number, 
       	cardholder_name,
       	expiration_date,
       	cvv_hash,
       	created_at
              FROM data_credit_cards
              WHERE data_credit_card_key = $1 AND private_user_key = $2`

	var dataCreditCard model.DataCreditCardResponse
	err := pstg.db.QueryRow(query, data.DataCreditCardKey, data.PrivateUserKey).Scan(
		&dataCreditCard.DataCreditCardKey,
		&dataCreditCard.CardNumber,
		&dataCreditCard.CardholderName,
		&dataCreditCard.ExpirationDate,
		&dataCreditCard.CVVHash,
		&dataCreditCard.CreatedAt,
	)

	if err != nil {
		return model.DataCreditCardResponse{}, err
	}

	return dataCreditCard, nil
}

func (pstg *PostgreSQL) DeleteDataCard(data model.DataCreditCard) error {
	query := `DELETE FROM data_credit_cards
              WHERE data_credit_card_key = $1 AND private_user_key = $2`

	_, err := pstg.db.Exec(query, data.DataCreditCardKey, data.PrivateUserKey)
	if err != nil {
		return err
	}

	return nil
}
