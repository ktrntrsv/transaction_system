CREATE TABLE account
(
    id      uuid PRIMARY KEY,
    balance float check ( balance >= 0 ) NOT NULL
);

INSERT INTO account(id, balance)
VALUES ('123e4567-e89b-12d3-a456-426614174000', 1000);

INSERT INTO account(id, balance)
VALUES ('123e4567-e89b-12d3-a456-426614174001', 1000);

INSERT INTO account(id, balance)
VALUES ('123e4567-e89b-12d3-a456-426614174002', 1000);

INSERT INTO account(id, balance)
VALUES ('123e4567-e89b-12d3-a456-426614174003', 1000);