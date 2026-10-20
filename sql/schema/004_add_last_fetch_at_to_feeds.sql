-- +goose Up
ALTER TABLE feeds 
ADD last_fetch_at TIMESTAMP;

-- +goose Down
ALTER TABLE feeds 
DROP last_fetch_at;