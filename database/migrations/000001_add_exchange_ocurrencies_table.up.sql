CREATE TABLE IF NOT EXISTS exchange_ocurrencies (
    id serial, 
    "from" text,
    "to" text,
    amount numeric,
    rate numeric,
    CONSTRAINT exchange_pkey PRIMARY KEY (id)
)