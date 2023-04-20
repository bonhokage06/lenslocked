-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions (user_id INT UNIQUE,remember_token text UNIQUE,CONSTRAINT fk_user
      FOREIGN KEY(user_id) 
	  REFERENCES users(id)
	  ON DELETE CASCADE,
	PRIMARY KEY(user_id)

);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;
-- +goose StatementEnd
