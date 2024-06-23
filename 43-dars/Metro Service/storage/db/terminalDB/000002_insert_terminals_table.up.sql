INSERT INTO Terminals (station_id)
VALUES
    ('aeba1901-dcd7-4dc4-9706-30f01276716b'),
    ('6a86a46b-7098-4b50-aee0-8194bf3031bb'),
    ('4cd1623e-d483-4455-b55a-1090177745ab'),
    ('60c184a4-2089-4fa8-af70-ba8cee2fd5da'),
    ('bc28406a-d39b-4d2a-8d84-096038f723c5'),
    ('6b33d0f7-edf7-4ade-94af-16029cc2c2c8'),
    ('d2d93600-0325-4503-80c5-5a070acd6e01'),
    ('378506ae-d8e9-43b5-bfec-1b46b7fda58e'),
    ('ab4bb698-8139-4fc1-b0e6-1621a703e31e'),
    ('421d57bc-2a45-42e7-bfb6-8d16997b4f4b'),
    ('97bed829-0b90-4fdb-9f7b-ed9f015308ab'),
    ('4e3eb30d-5060-4e7d-9317-d12d98a0ad3e'),
    ('f6046fb1-d1aa-4816-a852-3b052173450d'),
    ('17cc3a26-bf42-40c7-b9cd-debc1ac16ec7'),
    ('6463f700-be94-4ca7-b268-b510ddf6e9c0');



--migrate create -ext sql -dir storage/db/terminalDB -seq create_terminals_table
--migrate -database 'postgres://postgres:hamidjon4424@localhost:5432/metro?sslmode=disable' -path storage/db/terminalDB up