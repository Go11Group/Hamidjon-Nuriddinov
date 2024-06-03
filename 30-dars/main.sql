-- Users jadvali
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);

-- Products jadvali
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT NOT NULL
);


CREATE TABLE user_products(
    id SERIAL, 
    user_id int REFERENCES users(id), 
    product_id int REFERENCES products(id) 
);


INSERT INTO user_products(user_id, product_id) VALUES
(1, 1),
(1, 2),
(2, 3),
(2, 4),
(3, 5),
(3, 6),
(4, 7),
(4, 8),
(5, 9),
(5, 10),
(6, 11),
(7, 13),
(7, 14),
(8, 15),
(8, 1),
(9, 2),
(9, 3),
(10, 4),
(10, 5);

-- Crudlar:
-- Create (Yaratish):
-- CreateUser: Foydalanuvchi yaratish.
-- CreateProduct: Mahsulot yaratish.
-- Read (O'qish):
-- GetUser: Foydalanuvchini o'qish.
-- GetProduct: Mahsulotni o'qish.
-- Update (Yangilash):
-- UpdateUser: Foydalanuvchi ma'lumotlarini yangilash.
-- UpdateProduct: Mahsulot ma'lumotlarini yangilash.
-- Delete (O'chirish):
-- DeleteUser: Foydalanuvchini o'chirish.
-- DeleteProduct: Mahsulotni o'chirish.
-- user_products:(tableni o'zingiz yozing)
-- Bu tableda UserProduct ma'lumotlari saqlanadi. 
-- Aynan Userni productlarini create, read update qilishda transactionlardan foydalaning.


-- SAvollar bo'lsa so'rangiz
