CREATE TABLE users (
	id bigserial NOT NULL PRIMARY KEY,
	name varchar NOT NULL, 
	email varchar NOT NULL,
	password varchar NOT NULL,
	created_at timestamptz NOT NULL DEFAULT (now()) 
);