-- +migrate Up
CREATE TABLE IF NOT EXISTS users (id int NOT NULL AUTO_INCREMENT, email TEXT, password TEXT, activation Date, PRIMARY KEY (id) )

-- +migrate Down
DROP TABLE users;