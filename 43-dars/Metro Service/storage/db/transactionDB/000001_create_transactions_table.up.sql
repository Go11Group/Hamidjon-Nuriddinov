CREATE TYPE type as ENUM('credit', 'debit');

CREATE TABLE Transactions(
    transaction_id UUID default gen_random_uuid(),
    card_id UUID,
    amount real,
    terminal_id UUID,
    transaction_type type
);