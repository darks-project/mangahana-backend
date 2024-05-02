-- users
CREATE SCHEMA users;

CREATE TABLE users.users (
  id          SERIAL PRIMARY KEY,
  email       TEXT NOT NULL UNIQUE,
  password    TEXT NOT NULL,
  username    TEXT NOT NULL,
  description TEXT,
  photo       TEXT,
  role_id     INTEGER DEFAULT 1,
  is_banned   BOOLEAN NOT NULL DEFAULT FALSE,
  created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users.sessions (
  user_id    SERIAL NOT NULL,
  token      VARCHAR(128) NOT NULL UNIQUE,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users.roles (
  id          SERIAL PRIMARY KEY,
  name        TEXT NOT NULL,
  permissions INTEGER[] NOT NULL DEFAULT '{}'
);
INSERT INTO users.roles (name) VALUES ('Қолданушы');

CREATE TABLE users.permissions (
  id   SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);
INSERT INTO users.permissions (name) VALUES ('approve_team');




-- teams
CREATE SCHEMA teams;

CREATE TABLE teams.teams (
  id          SERIAL PRIMARY KEY,
  owner_id    SERIAL NOT NULL,
  type_id     SERIAL NOT NULL,
  name        VARCHAR(64) NOT NULL,
  description TEXT,
  photo       TEXT,
  approved    BOOLEAN NOT NULL DEFAULT FALSE,
  created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE teams.types (
  id   SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);
INSERT INTO teams.types (name) VALUES ('Аударма'), ('Баспа'), ('Автор');

CREATE TABLE teams.members (
  team_id SERIAL NOT NULL,
  user_id SERIAL NOT NULL
);