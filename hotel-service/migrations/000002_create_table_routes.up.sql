-- +migrate Up
CREATE TABLE routes (
    id SERIAL PRIMARY KEY,
    vehicle_id INT NOT NULL,
    src VARCHAR(255) NOT NULL,
    dest VARCHAR(255) NOT NULL,
    catagory VARCHAR(255) NOT NULL,
    cost VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE
);



