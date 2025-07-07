-- +migrate Up
CREATE TABLE pictures (
    id SERIAL PRIMARY KEY,
    vehicle_id INT NOT NULL,
    urls TEXT[]
);