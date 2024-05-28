CREATE TABLE "branch" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "title" VARCHAR(30) NOT NULL
);

CREATE TABLE "teacher" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "branch_id" UUID REFERENCES "branch" ("id"),
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "salary" NUMERIC
);

CREATE TABLE "asisstant_teacher" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "branch_id" UUID REFERENCES "branch" ("id"),
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "salary" NUMERIC
);


CREATE TABLE "course" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "name" VARCHAR(30) NOT NULL,
    "teacher_id" UUID REFERENCES "teacher" ("id"),
    "asisstant_teacher_id" UUID REFERENCES "asisstant_teacher" ("id"),
    "price" NUMERIC
);

CREATE TABLE "student" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "course_id" UUID REFERENCES "course" ("id"),
    "payed_sum" NUMERIC
);

INSERT INTO "branch" ("title") VALUES
('Chilonzor'),
('Hadra'),
('Beruniy'),
('Farg''ona'),
('Xorazm');


INSERT INTO teacher(branch_id, first_name, last_name) VALUES
('d1640946-8f83-4ba9-ae30-3cee4d1913f9', 'Nodirbek', 'Qobulov'),
('d1640946-8f83-4ba9-ae30-3cee4d1913f9', 'Yusufxo''ja', 'Turg''unxo''jayev'),
('d1640946-8f83-4ba9-ae30-3cee4d1913f9', 'Akbarshox', 'Sattorov'),
('d1640946-8f83-4ba9-ae30-3cee4d1913f9', 'Husan', 'Musayev');

INSERT INTO asisstant_teacher(branch_id, first_name, last_name) VALUES
('d1640946-8f83-4ba9-ae30-3cee4d1913f9', 'Muhammadali', 'Zoirov'),
('d1640946-8f83-4ba9-ae30-3cee4d1913f9', 'Anvarjon', 'Shavqiyev'),
('d1640946-8f83-4ba9-ae30-3cee4d1913f9', 'Zafar', 'Samadov');

INSERT INTO course(name, teacher_id, asisstant_teacher_id, price) VALUES
('Foundation', 'e92264c3-65f3-46a6-a71f-2ddac539eb52', 'f2e9f12d-3757-4055-97d8-1083f42d9e3f', 2200000),
('Foundation', '866f7860-b9f6-4a4d-9427-4f1d3546ccad', 'f2e9f12d-3757-4055-97d8-1083f42d9e3f', 2200000),
('Foundation', '5b1c86dc-18a1-4f16-a15d-03a661142341', '29ab50bf-04aa-40f2-b2be-c55efaaf71b8', 2200000),
('GO', 'fda9a1cf-225b-45da-afdc-36f5a80c684a', '07e97d95-59bf-4b5d-8e49-749ff55028ae', 2200000);

INSERT INTO "student" ("first_name", "last_name", "course_id", "payed_sum") VALUES
('John', 'Doe', '5022cd64-47af-4d67-a3be-eb6c669c0c87', 2200000),
('Jane', 'Doe', '5b9a5ad4-5df7-4823-a536-0a1cb0793bf7', 2200000),
('Alice', 'Smith', '37651ae5-9a36-4cd6-844f-60c02c52d158', 2200000),
('Bob', 'Brown', '79ae88a8-ef5e-4888-b6aa-b4676598013f', 2200000),
('Charlie', 'Johnson', '5022cd64-47af-4d67-a3be-eb6c669c0c87', 2200000),
('David', 'Williams', '5b9a5ad4-5df7-4823-a536-0a1cb0793bf7', 2200000),
('Eve', 'Jones', '37651ae5-9a36-4cd6-844f-60c02c52d158', 2200000),
('Frank', 'Miller', '79ae88a8-ef5e-4888-b6aa-b4676598013f', 2200000),
('Grace', 'Davis', '5022cd64-47af-4d67-a3be-eb6c669c0c87', 2200000),
('Hank', 'Garcia', '5b9a5ad4-5df7-4823-a536-0a1cb0793bf7', 2200000),
('Ivy', 'Martinez', '37651ae5-9a36-4cd6-844f-60c02c52d158', 2200000),
('Jack', 'Rodriguez', '79ae88a8-ef5e-4888-b6aa-b4676598013f', 2200000),
('Karen', 'Martinez', '5022cd64-47af-4d67-a3be-eb6c669c0c87', 2200000),
('Leo', 'Anderson', '5b9a5ad4-5df7-4823-a536-0a1cb0793bf7', 2200000),
('Mona', 'Thomas', '37651ae5-9a36-4cd6-844f-60c02c52d158', 2200000),
('Nina', 'Jackson', '79ae88a8-ef5e-4888-b6aa-b4676598013f', 2200000),
('Oscar', 'White', '5022cd64-47af-4d67-a3be-eb6c669c0c87', 2200000),
('Pam', 'Harris', '5b9a5ad4-5df7-4823-a536-0a1cb0793bf7', 2200000),
('Quinn', 'Clark', '37651ae5-9a36-4cd6-844f-60c02c52d158', 2200000),
('Rita', 'Lewis', '79ae88a8-ef5e-4888-b6aa-b4676598013f', 2200000),
('Steve', 'Adams', '5022cd64-47af-4d67-a3be-eb6c669c0c87', 2100000),
('Tina', 'Baker', '5b9a5ad4-5df7-4823-a536-0a1cb0793bf7', 1900000),
('Uma', 'Clark', '37651ae5-9a36-4cd6-844f-60c02c52d158', 1800000),
('Victor', 'Davis', '79ae88a8-ef5e-4888-b6aa-b4676598013f', 1700000),
('Wendy', 'Evans', '5022cd64-47af-4d67-a3be-eb6c669c0c87', 1600000),
('Xander', 'Garcia', '5b9a5ad4-5df7-4823-a536-0a1cb0793bf7', 1500000),
('Yara', 'Harris', '37651ae5-9a36-4cd6-844f-60c02c52d158', 1400000),
('Zane', 'Johnson', '79ae88a8-ef5e-4888-b6aa-b4676598013f', 1300000),
('Alex', 'King', '5022cd64-47af-4d67-a3be-eb6c669c0c87', 1200000),
('Bella', 'Lewis', '5b9a5ad4-5df7-4823-a536-0a1cb0793bf7', 1100000);



select sum(C.price)
from student as S
join course as C 
on S.course_id = C.id
join teacher as T 
on C.teacher_id = T.id
group by C.teacher_id, T.first_name, T.last_name;


--1
Update teacher set salary = (
    select sum(C.price) * 0.5
from course as C
join teacher as T 
on C.teacher_id = T.id
where teacher.id = T.id
group by C.teacher_id
);

--1.1
with NewTable as (
    select sum(C.price) * 0.2 as salary, C.asisstant_teacher_id as at_id
from course as C 
join asisstant_teacher as At 
on C.asisstant_teacher_id = At.id
group by C.asisstant_teacher_id
)

update asisstant_teacher 
set salary = NT.salary
from NewTable as NT
where id = NT.at_id;


--2 
select T.first_name, T.last_name, array_agg(distinct C.name), array_agg(S.last_name)
from teacher as T
join course as C 
on C.teacher_id = T.id
join student as S 
on S.course_id = C.id 
group by T.id;

--3
select S.first_name, S.last_name, S.payed_sum as 'to''lagan puli', C.price-S.payed_sum as 'qarzdorlik'
from student as S
join course as C
on S.course_id = C.id
where S.payed_sum < C.price;

