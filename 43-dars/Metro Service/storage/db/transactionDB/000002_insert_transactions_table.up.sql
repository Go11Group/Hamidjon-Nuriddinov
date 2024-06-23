INSERT INTO Transactions (card_id, amount, terminal_id, transaction_type)
VALUES
    ('40a02c66-9100-4841-850d-1c51c4d7bfe7', 30050.0, '6a86a46b-7098-4b50-aee0-8194bf3031bb', 'debit'),
    ('468726c4-dbf0-410a-ba2c-9502cf4885e1', 7500.0, '60c184a4-2089-4fa8-af70-ba8cee2fd5da', 'debit'),
    ('8b2a17bd-0d5f-4bc8-ba2a-1a36c4f0db25', 15025.0, '6b33d0f7-edf7-4ade-94af-16029cc2c2c8', 'debit'),
    ('6d19aa92-71f1-4b92-a4dc-5a5d95c86a1f', 10075.0, '378506ae-d8e9-43b5-bfec-1b46b7fda58e', 'debit'),
    ('49ee7925-f3da-47c8-ae9f-5672eaef48a5', 9025.0, '421d57bc-2a45-42e7-bfb6-8d16997b4f4b', 'debit'),
    ('91dcf567-ea21-439a-930f-93332e20d52f', 30050.0, '4e3eb30d-5060-4e7d-9317-d12d98a0ad3e', 'debit'),
    ('73a1c37e-4aec-4697-b3ab-ca168b597088', 7500.0, '17cc3a26-bf42-40c7-b9cd-debc1ac16ec7', 'debit');



--migrate create -ext sql -dir storage/db/transactionDB -seq create_transactions_table
--migrate -database 'postgres://postgres:hamidjon4424@localhost:5432/metro?sslmode=disable' -path storage/db/transactionDB up