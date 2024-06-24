INSERT INTO Transactions (card_id, amount, terminal_id, transaction_type)
VALUES
    ('5b34f33c-ff68-47a4-ac92-d77acb32ca6c', 30050.0, 'aeba1901-dcd7-4dc4-9706-30f01276716b', 'debit'),
    ('e0dec45d-0338-4f71-b4aa-1c9eba006f55', 7500.0, '6a86a46b-7098-4b50-aee0-8194bf3031bb', 'debit'),
    ('224177d0-9a28-4956-b39c-2a50f659b366', 15025.0, '4cd1623e-d483-4455-b55a-1090177745ab', 'debit'),
    ('2db53d2c-52f0-4b3e-bccc-bd6e9c27d60f', 10075.0, 'd2d93600-0325-4503-80c5-5a070acd6e01', 'debit'),
    ('3e25910c-fd56-4ca6-9dac-94d7da0a134f', 9025.0, '421d57bc-2a45-42e7-bfb6-8d16997b4f4b', 'debit'),
    ('032b1c57-38c9-4b33-9a15-6c79fe40a0c8', 30050.0, '421d57bc-2a45-42e7-bfb6-8d16997b4f4b', 'debit'),
    ('fc86a901-8f3d-45f5-8bc9-38cdc06cb2a7', 7500.0, 'bc28406a-d39b-4d2a-8d84-096038f723c5', 'debit');



--migrate create -ext sql -dir storage/db/transactionDB -seq create_transactions_table
--migrate -database 'postgres://postgres:hamidjon4424@localhost:5432/metro?sslmode=disable' -path storage/db/transactionDB up



