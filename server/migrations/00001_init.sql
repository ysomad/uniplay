-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS demos (
    id uuid PRIMARY KEY NOT NULL,
    status varchar(16) NOT NULL,
    reason varchar(255),
    identity_id uuid NOT NULL,
    uploaded_at timestamptz NOT NULL,
    processed_at timestamptz
);
-- +goose StatementEnd
