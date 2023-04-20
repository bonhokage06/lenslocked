-- +goose Up
-- +goose StatementBegin
CREATE TABLE widgets (id INT GENERATED ALWAYS AS IDENTITY, color text not null,PRIMARY KEY(id));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE widgets;
-- +goose StatementEnd
