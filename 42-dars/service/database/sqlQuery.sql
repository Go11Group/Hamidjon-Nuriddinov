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



INSERT INTO Users (name, email, birthday, password)
VALUES
('John Doe', 'john.doe@example.com', '1990-01-01', 'password1'),
('Jane Smith', 'jane.smith@example.com', '1992-02-02', 'password2'),
('Alice Johnson', 'alice.johnson@example.com', '1994-03-03', 'password3'),
('Bob Brown', 'bob.brown@example.com', '1988-04-04', 'password4'),
('Charlie Davis', 'charlie.davis@example.com', '1985-05-05', 'password5'),
('Emily Wilson', 'emily.wilson@example.com', '1991-06-06', 'password6'),
('Frank Miller', 'frank.miller@example.com', '1989-07-07', 'password7'),
('Grace Lee', 'grace.lee@example.com', '1993-08-08', 'password8'),
('Hank Young', 'hank.young@example.com', '1995-09-09', 'password9'),
('Ivy Hall', 'ivy.hall@example.com', '1986-10-10', 'password10'),
('Jack King', 'jack.king@example.com', '1987-11-11', 'password11'),
('Kate Wright', 'kate.wright@example.com', '1990-12-12', 'password12'),
('Leo Scott', 'leo.scott@example.com', '1992-01-13', 'password13'),
('Mia Green', 'mia.green@example.com', '1994-02-14', 'password14'),
('Nate Adams', 'nate.adams@example.com', '1988-03-15', 'password15'),
('Olivia Clark', 'olivia.clark@example.com', '1991-04-16', 'password16'),
('Paul Turner', 'paul.turner@example.com', '1993-05-17', 'password17'),
('Quinn Hill', 'quinn.hill@example.com', '1985-06-18', 'password18'),
('Rachel Allen', 'rachel.allen@example.com', '1987-07-19', 'password19'),
('Steve Baker', 'steve.baker@example.com', '1989-08-20', 'password20');


INSERT INTO Courses (title, description)
VALUES
('Introduction to Programming', 'Learn the basics of programming using Python.'),
('Advanced Python', 'Deep dive into advanced Python topics and techniques.'),
('Web Development with Django', 'Build web applications using the Django framework.'),
('Data Science with Python', 'Introduction to data science concepts using Python.'),
('Machine Learning', 'An overview of machine learning algorithms and applications.'),
('Database Design', 'Learn how to design and implement relational databases.'),
('Cloud Computing', 'Introduction to cloud computing and its applications.'),
('Mobile App Development', 'Create mobile applications for Android and iOS.'),
('Cybersecurity Fundamentals', 'Learn the basics of cybersecurity and how to protect systems.'),
('DevOps Essentials', 'Understand the principles and practices of DevOps.');


INSERT INTO Lessons (course_id, title, content)
VALUES
('fc2987ca-241c-4074-bec1-2ec18500c64a', 'Lesson 1: Introduction to Programming', 'This lesson covers the basics of programming.'),
('0b7fbb5f-4193-497d-a0b6-fef79640159a', 'Lesson 2: Variables and Data Types', 'This lesson explains variables and data types.'),
('e575bba1-280c-447d-a791-e87e9a7d536e', 'Lesson 1: Advanced Functions', 'This lesson dives into advanced functions in Python.'),
('dea3ed91-c2da-44e4-9089-327e8e938b10', 'Lesson 2: Decorators and Generators', 'This lesson covers decorators and generators in Python.'),
('ed5a1bc4-32b5-417b-9f79-a9c6abc9e8b6', 'Lesson 1: Setting Up Django', 'This lesson guides you through setting up Django.'),
('30ff4ac5-47d6-4c76-b627-f4e5b9af8e20', 'Lesson 2: Django Models', 'This lesson explains how to create models in Django.'),
('0314a636-2d1c-485d-98a4-44558c6186fa', 'Lesson 1: Introduction to Data Science', 'This lesson provides an overview of data science.'),
('dea3ed91-c2da-44e4-9089-327e8e938b10', 'Lesson 2: Data Analysis with Pandas', 'This lesson covers data analysis using the Pandas library.'),
('7f8f5db1-91cf-4594-9845-634ab27ba61f', 'Lesson 1: Introduction to Machine Learning', 'This lesson introduces machine learning concepts.'),
('de4e2af1-70dc-42bc-81c3-3fcf23e21671', 'Lesson 2: Supervised Learning', 'This lesson covers supervised learning techniques.');


INSERT INTO Enrollments (user_id, course_id)
VALUES
('68ec648a-cb75-4fdc-bdea-6680fcd8eaab', 'fc2987ca-241c-4074-bec1-2ec18500c64a'),
('68ec648a-cb75-4fdc-bdea-6680fcd8eaab', 'd1579eae-a587-4d23-adcf-3923b8699162'),
('d876de02-7d63-4e6e-9ebb-71925e6a7222', 'de4e2af1-70dc-42bc-81c3-3fcf23e21671'),
('7474a0c6-bea2-4f7c-8d94-6cbb568b655c', '0314a636-2d1c-485d-98a4-44558c6186fa'),
('0d078bbc-7557-4be4-8544-f836d546db0e', '2b32ae32-698d-4305-95ce-8d0e5012482d'),
('d6cb761e-dbc3-47a7-90e5-4b8ba2126bcf', 'fb749e63-8b2f-4a33-a965-573014fb4bf1'),
('eaf8f40b-a8ca-4764-aff7-90ac763333db', 'dea3ed91-c2da-44e4-9089-327e8e938b10'),
('479c18ce-e524-40ae-aa5c-6945430b4127', 'ed5a1bc4-32b5-417b-9f79-a9c6abc9e8b6'),
('f12ab8a3-3a64-4394-9e5c-4d1deddb7b37', '2b32ae32-698d-4305-95ce-8d0e5012482d'),
('0ed31c06-e7ae-4abf-a0aa-4c739966c0f0', '0b7fbb5f-4193-497d-a0b6-fef79640159a');


 