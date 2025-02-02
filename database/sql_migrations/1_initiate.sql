-- +migrate Up
-- +migrate StatementBegin

create table anime (
    id BIGINT NOT NULL,
    title varchar(255),
    description_anime TEXT,
    genre varchar(255),
    release_year INTEGER
)

-- +migrate StatementEnd