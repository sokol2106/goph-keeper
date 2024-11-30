CREATE database gophkeeper;
comment ON database gophkeeper IS 'Менеджер паролей';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE private_user
(
    private_user_key uuid DEFAULT uuid_generate_v4() NOT NULL
        CONSTRAINT private_user_pk
            PRIMARY KEY
        CONSTRAINT private_user_private_user_private_user_key_fk
            REFERENCES private_user
            ON UPDATE CASCADE
            ON DELETE CASCADE
            DEFERRABLE INITIALLY DEFERRED,
    login text,
    password_hash text,
    encryption_key text
);

COMMENT ON TABLE private_user IS 'Пользователи';

COMMENT ON COLUMN private_user.encryption_key IS 'Ключ шифрования';

CREATE TABLE data_text
(
    data_text_key uuid DEFAULT uuid_generate_v4() NOT NULL
        CONSTRAINT data_text_pk
            PRIMARY KEY,
    private_user_key uuid NOT NULL
        CONSTRAINT data_text_private_user_private_user_key_fk
            REFERENCES private_user
            ON UPDATE CASCADE
            ON DELETE CASCADE
            DEFERRABLE INITIALLY DEFERRED,
    data text
);

COMMENT ON COLUMN data_text.data IS 'Произвольные текстовые данные';

CREATE TABLE data_credit_cards
(
    data_credit_card_key uuid DEFAULT uuid_generate_v4() NOT NULL
        CONSTRAINT data_credit_cards_pk
            PRIMARY KEY,
    card_number text NOT NULL,
    cardholder_name text NOT NULL,
    expiration_date text NOT NULL,
    cvv_hash text NOT NULL,
    created_at date DEFAULT now() NOT NULL,
    private_user_key uuid NOT NULL
);

COMMENT ON COLUMN data_credit_cards.card_number IS 'Номер карты (в зашифрованном виде)';

COMMENT ON COLUMN data_credit_cards.cardholder_name IS 'Имя владельца карты';

COMMENT ON COLUMN data_credit_cards.expiration_date IS 'Срок действия карты';

COMMENT ON COLUMN data_credit_cards.cvv_hash IS 'Хэш CVV-кода';

COMMENT ON COLUMN data_credit_cards.created_at IS 'Дата добавления записи';

CREATE TABLE data_binary
(
    data_binary_key uuid DEFAULT uuid_generate_v4() NOT NULL
        CONSTRAINT data_binary_pk
            PRIMARY KEY,
    private_user_key uuid NOT NULL,
    filename text NOT NULL,
    data bytea NOT NULL
);

COMMENT ON TABLE data_binary IS 'Произвольные бинарные данные';

COMMENT ON COLUMN data_binary.filename IS 'Наименование файла';

COMMENT ON COLUMN data_binary.data IS 'Данные';

