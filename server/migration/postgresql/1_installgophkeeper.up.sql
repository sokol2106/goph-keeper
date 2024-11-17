CREATE database gophkeeper;
comment ON database gophkeeper IS 'Менеджер паролей';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.users
(
    user_key uuid DEFAULT uuid_generate_v4() NOT NULL
        CONSTRAINT users_pk
            PRIMARY KEY,
    login    text,
    password text
);

COMMENT ON TABLE public.users IS 'Пользователи';

