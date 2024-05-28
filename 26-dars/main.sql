create table Students(student_id serial primary key, name varchar, age int);

create table Courses(course_id serial primary key, name varchar);

create table Student_course(
    student_course_id serial primary key,
    student_id int, 
    course_id int,
    foreign key(student_id) references Students(student_id),
    foreign key(course_id) references Courses(course_id));

create table Grade(grade_id serial, grade int, student_course_id int, course_id int,
    foreign key(student_course_id) references Student_course(student_course_id));


INSERT INTO Students (name, age) VALUES 
('Azizbek', 20),
('Dilshod', 22),
('Shirin', 21),
('Jasur', 23),
('Otabek', 19),
('Gulnora', 20),
('Kamola', 22),
('Farhod', 21),
('Lola', 23),
('Bekzod', 19),
('Zarina', 20),
('Umid', 22),
('Shahzoda', 21),
('Alisher', 23),
('Nilufar', 19),
('Firdavs', 20),
('Dilafruz', 22),
('Sherzod', 21),
('Sardor', 23),
('Yulduz', 19);

INSERT INTO Courses (name) VALUES 
('Matematika'),
('Fizika'),
('Kimyo'),
('Biologiya'),
('Ingliz tili'),
('Tarix'),
('Geografiya'),
('Informatika'),
('Falsafa'),
('Adabiyot');


INSERT INTO Student_course (student_id, course_id) VALUES 
(1, 1),
(1, 2),
(2, 1),
(2, 3),
(3, 2),
(3, 4),
(4, 1),
(4, 5),
(5, 3),
(5, 6),
(6, 2),
(6, 7),
(7, 1),
(7, 8),
(8, 4),
(8, 9),
(9, 3),
(9, 10),
(10, 5),
(10, 1),
(11, 6),
(11, 2),
(12, 7),
(12, 3),
(13, 8),
(13, 4),
(14, 9),
(14, 5),
(15, 10),
(15, 6),
(16, 1),
(16, 7),
(17, 2),
(17, 8),
(18, 3),
(18, 9),
(19, 4),
(19, 10),
(20, 5),
(20, 1);

INSERT INTO Grade (grade, student_course_id, course_id) VALUES 
(85, 1, 1),
(90, 2, 2),
(75, 3, 1),
(80, 4, 3),
(88, 5, 2),
(92, 6, 4),
(78, 7, 1),
(85, 8, 5),
(70, 9, 3),
(95, 10, 6),
(82, 11, 2),
(87, 12, 7),
(91, 13, 1),
(84, 14, 8),
(73, 15, 4),
(89, 16, 9),
(76, 17, 3),
(94, 18, 10),
(81, 19, 5),
(88, 20, 1),
(77, 21, 6),
(92, 22, 2),
(83, 23, 7),
(79, 24, 3),
(90, 25, 8),
(86, 26, 4),
(75, 27, 9),
(91, 28, 5),
(82, 29, 10),
(88, 30, 6),
(84, 31, 1),
(89, 32, 7),
(78, 33, 2),
(92, 34, 8),
(85, 35, 3),
(87, 36, 9),
(80, 37, 4),
(94, 38, 10),
(79, 39, 5),
(90, 40, 1);


--3
select C.name, round(avg(G.grade)::numeric, 2)
from Courses as C 
join Student_course as Sc 
on C.course_id = Sc.course_id
join Grade as G 
on Sc.student_course_id = G.student_course_id
group by C.name;

--4 
select C.name, array_agg(S.name) as "Yosh o'quvchilar", min(S.age) as "Minimal yosh"
from Students as S
join Student_course as Sc 
on S.student_id = Sc.student_id
join Courses as C 
on C.course_id = Sc.course_id
group by C.name;

--5
select C.name, round(avg(G.grade)::numeric, 2) as grade 
from Courses as C 
join Student_course as Sc 
on C.course_id = Sc.course_id
join Grade as G 
on Sc.student_course_id = G.student_course_id
where grade = (
    select round(avg(G.grade)::numeric, 2) as grade
    from Courses as C
    join Student_course as Sc
    on C.course_id = Sc.course_id
    join Grade as G
    on G.student_course_id = Sc.student_course_id
    group by C.name
    order by grade desc limit 1
)
group by C.name;

--2
select C.name, array_agg(S.name), G.grade
from Students as S
join Student_course as Sc
on Sc.student_id = S.student_id
join Courses as C
on C.course_id = Sc.course_id
join Grade as G 
on G.student_course_id = Sc.student_course_id
join (
    select C.name as name, max(G.grade) as grade
from Students as S
join Student_course as Sc
on S.student_id = Sc.student_id
join Courses as C 
on C.course_id = Sc.course_id
join Grade as G
on Sc.student_course_id = G.student_course_id
group by C.name
) as NewTable
on C.name = NewTable.name and G.grade = NewTable.grade
group by C.name, G.grade;