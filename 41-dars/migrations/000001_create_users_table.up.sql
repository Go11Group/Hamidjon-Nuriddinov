CREATE TABLE Users(
    id uuid default gen_random_uuid(),
    first_name varchar,
    last_name varchar,
    age int,
    nation varchar
)

