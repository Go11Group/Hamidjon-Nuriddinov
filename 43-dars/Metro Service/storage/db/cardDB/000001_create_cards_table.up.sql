CREATE TABLE Cards(
    card_id UUID default gen_random_uuid(),
    number varchar,
    user_id UUID
);