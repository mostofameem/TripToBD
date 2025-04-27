-- +migrate Up
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    vehicle_id INT NOT NULL,
    review TEXT NOT NULL,
    reviewer_id INT NOT NULL,
    reviewed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE
);