CREATE TABLE  IF NOT EXISTS users (
    id uuid default gen_random_uuid() primary key, 
    email varchar not null,
    password text not null,
    activated boolean default false,
    activated_at date,
    username varchar(50) not null,
    roles text[],
    created_at timestamptz default CURRENT_TIMESTAMP not null,
    updated_at timestamptz default CURRENT_TIMESTAMP not null,
    UNIQUE(email)
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users 
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();