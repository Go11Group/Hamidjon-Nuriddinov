CREATE DATABASE Project;


CREATE TABLE Users(
    user_id uuid default gen_random_uuid() Primary Key,          -- Foydalanuvchi ID
    name varchar(100),                                           -- Foydalanuvchining ismi
    email varchar(100),                                          -- Foydalanuvchining elektron pochtasi
    birthday timestamp,                                          -- Foydalanuvchi tug’ilgan vaqti
    password varchar(100),                                       -- Foydalanuvchining paroli
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,              -- Yaratilgan vaqti
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,              -- Yangilangan vaqti
    deleted_at timestamp                                         -- O'chirilgan vaqti
);


CREATE TABLE Courses(
    course_id uuid default gen_random_uuid() Primary Key,        -- Kurs ID
    title varchar(100),                                          -- Kurs nomi
    description text,                                            -- Kurs tavsifi
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,              -- Yaratilgan vaqti
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,              -- Yangilangan vaqti
    deleted_at timestamp                                         -- O'chirilgan vaqti
);


CREATE TABLE Lessons(
    lesson_id uuid default gen_random_uuid() Primary Key,        -- Dars ID
    course_id uuid REFERENCES Courses(course_id),                -- Kurs ID (Foreign Key)
    title varchar(100),                                          -- Dars nomi
    content text,                                                -- Dars kontenti
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,              -- Yaratilgan vaqti
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,              -- Yangilangan vaqti
    deleted_at timestamp                                         -- O'chirilgan vaqti
);


CREATE TABLE Enrollments(
    enrollment_id uuid default gen_random_uuid() Primary Key,    -- Ro'yxatdan o'tish ID
    user_id uuid REFERENCES Users(user_id),                      -- Foydalanuvchi ID (Foreign Key)
    course_id uuid REFERENCES Courses(course_id),                -- Kurs ID (Foreign Key)
    enrollment_date timestamp DEFAULT CURRENT_TIMESTAMP,         -- Ro'yxatdan o'tish sanasi
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,              -- Yaratilgan vaqti
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,              -- Yangilangan vaqti
    deleted_at timestamp                                         -- O'chirilgan vaqti
);