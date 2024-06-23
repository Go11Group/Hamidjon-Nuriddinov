CREATE TABLE Users(
    user_id UUID default gen_random_uuid(),
    name varchar,
    phone varchar,
    age int 
)