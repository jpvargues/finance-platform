CREATE TABLE etfs
(
    id SERIAL PRIMARY KEY,
    ticker varchar(255) UNIQUE,
    name varchar(255),
    isin CHAR(12) UNIQUE,
    is_accumulating BOOLEAN,
    category varchar(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE etf_holdings
(
    id SERIAL PRIMARY KEY,
    etf_id INT REFERENCES etfs(id),
    company_name varchar(255),
    percent NUMERIC(5,2),
    as_of_date DATE DEFAULT CURRENT_DATE
);

CREATE TABLE etf_market_exposure
(
    id SERIAL PRIMARY KEY,
    etf_id INT REFERENCES etfs(id),
    market varchar(255),
    percent NUMERIC(5,2),
    as_of_date DATE DEFAULT CURRENT_DATE
);

CREATE TABLE etf_sector_exposure
(
    id SERIAL PRIMARY KEY,
    etf_id INT REFERENCES etfs(id),
    sector varchar(255),
    percent NUMERIC(5,2),
    as_of_date DATE DEFAULT CURRENT_DATE
);
