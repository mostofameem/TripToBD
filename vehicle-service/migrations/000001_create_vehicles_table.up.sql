-- +migrate Up
CREATE TABLE vehicles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    catagory VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    picture_url TEXT NOT NULL,
    rating REAL DEFAULT 0.0,
    vote_count INT DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);
