CREATE TABLE IF NOT EXISTS price (
    id SERIAL PRIMARY KEY,
    coin TEXT,
    price NUMERIC,
    "timestamp" INTEGER
);