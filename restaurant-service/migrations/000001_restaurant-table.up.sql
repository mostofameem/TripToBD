-- +migrate Up
CREATE TABLE IF NOT EXISTS restaurants (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    picture_url TEXT,
    location_info VARCHAR(255),
    rating DOUBLE PRECISION,
    vote_count INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    is_active BOOLEAN
);
