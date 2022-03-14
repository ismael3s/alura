CREATE TABLE personalities (
    id SERIAL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);

ALTER TABLE personalities ADD CONSTRAINT personalities_pk PRIMARY KEY (id);

INSERT INTO personalities(name, description) VALUES ('John', 'John is a good guy');
INSERT INTO personalities(name, description) VALUES ('Doe', 'Doe is a nbad guy');