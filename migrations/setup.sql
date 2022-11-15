CREATE TABLE account
(
    id      uuid PRIMARY KEY,
    balance float check ( balance >= 0 ) NOT NULL
);

INSERT INTO account(id, balance)
VALUES ('123e4567-e89b-12d3-a456-426614174000', 0);

INSERT INTO account(id, balance)
VALUES ('123e4567-e89b-12d3-a456-426614174001', 0);

INSERT INTO account(id, balance)
VALUES ('123e4567-e89b-12d3-a456-426614174002', 0)