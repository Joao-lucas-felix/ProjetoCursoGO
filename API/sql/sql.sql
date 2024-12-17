-- Criação do banco (não há IF NOT EXISTS no PostgreSQL)
-- Este comando precisa ser executado separadamente
CREATE DATABASE devbook;

-- Remoção da tabela, se já existir
DROP TABLE IF EXISTS usuarios;

-- Criação da tabela
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    name VARCHAR(250) NOT NULL,
    nick VARCHAR(250) NOT NULL UNIQUE,
    email VARCHAR(250) NOT NULL UNIQUE,
    password_hash VARCHAR(500) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON usuarios
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
