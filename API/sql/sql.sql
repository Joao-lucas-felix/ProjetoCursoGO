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


CREATE TABLE seguidores (
    usuario_id INT NOT NULL,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE, 

    seguidor_id INT NOT NULL,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE, 

    PRIMARY KEY (usuario_id, seguidor_id)
);

INSERT INTO usuarios (name, nick, email, password_hash) VALUES
    ('Usuario 1', 'Usuario 1', 'usuario1@gmail.com', '$2a$10$/BNFEUIM0spGEH.R9tjYJOgHd2npOvLC7vdxFi5MjlBjws5/H1Yxu'),
    ('Usuario 2', 'Usuario 2', 'usuario2@gmail.com', '$2a$10$/BNFEUIM0spGEH.R9tjYJOgHd2npOvLC7vdxFi5MjlBjws5/H1Yxu'),
    ('Usuario 3', 'Usuario 3', 'usuario3@gmail.com', '$2a$10$/BNFEUIM0spGEH.R9tjYJOgHd2npOvLC7vdxFi5MjlBjws5/H1Yxu');

INSERT INTO seguidores (usuario_id, seguidor_id) VALUES 
    (1,2),
    (1,3),
    (3,1);