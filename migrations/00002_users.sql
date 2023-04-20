-- +goose Up
-- +goose StatementBegin
CREATE TABLE users 
(id INT GENERATED ALWAYS AS IDENTITY,
email text unique not null, 
password_hash text not null,PRIMARY KEY(id));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
