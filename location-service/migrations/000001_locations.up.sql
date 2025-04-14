-- +migrate Up
CREATE TABLE IF NOT EXISTS locations (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content VARCHAR(255),
    best_time VARCHAR(50),
    picture_url TEXT,
    rating DOUBLE PRECISION,
    vote_count INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    is_active BOOLEAN
);