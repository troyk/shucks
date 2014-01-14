-- +goose Up

CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE users (
  id serial primary key,
  created_at timestamptz not null default now(),
  username citext not null UNIQUE CHECK (username ~ '^[a-zA-Z]{1}[a-zA-Z_0-9]{0,23}$'),
  email citext not null UNIQUE CHECK (length(email) < 180),
  mobile_phone citext UNIQUE,
  name citext,
  password bytea
);

