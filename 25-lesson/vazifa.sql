-- Many to manyga misol   
-- student va course jadvallari bilan qildim.


create table Students(student_id serial primary key, FISH varchar, age int, scolarship real);

create table Courses(course_id serial primary key, name varchar, teacher varchar);

create table Student_course(
    student_id int, course_id int, 
    foreign key (student_id) references Students(student_id), 
    foreign key (course_id) references Courses(course_id));


INSERT INTO Students (FISH, age, scolarship) VALUES
('Alice Johnson', 20, 1500.50),
('Bob Smith', 22, 2000.00),
('Charlie Brown', 21, 1750.75),
('David Williams', 23, 1800.25),
('Eve Davis', 19, 1000.00),
('Frank Miller', 20, 1600.50),
('Grace Wilson', 21, 1700.00),
('Hank Moore', 22, 1950.00),
('Ivy Taylor', 23, 2000.50),
('Jack Anderson', 19, 1100.75),
('Kara Thomas', 20, 1200.25),
('Leo Martinez', 21, 1400.00),
('Mia Robinson', 22, 1500.00),
('Nina Clark', 23, 1900.50),
('Oscar Lewis', 19, 1300.75),
('Paula Walker', 20, 1250.00),
('Quinn Hall', 21, 1350.50),
('Rick Young', 22, 1450.00),
('Sophie King', 23, 1550.25),
('Tom Scott', 19, 1050.75);


INSERT INTO Courses (name, teacher) VALUES
('Introduction to Computer Science', 'Dr. Alice Johnson'),
('Calculus I', 'Prof. Bob Smith'),
('Physics I', 'Dr. Charlie Brown'),
('Chemistry I', 'Dr. David Williams'),
('English Literature', 'Prof. Eve Davis'),
('World History', 'Dr. Frank Miller'),
('Biology I', 'Prof. Grace Wilson'),
('Introduction to Psychology', 'Dr. Hank Moore'),
('Sociology', 'Prof. Ivy Taylor'),
('Economics I', 'Dr. Jack Anderson');


INSERT INTO Student_course (student_id, course_id) VALUES
(1, 1),
(1, 2),
(2, 1),
(2, 3),
(3, 2),
(3, 4),
(4, 3),
(4, 5),
(5, 1),
(5, 6),
(6, 2),
(6, 7),
(7, 3),
(7, 8),
(8, 4),
(8, 9),
(9, 5),
(9, 10),
(10, 1),
(10, 2),
(11, 3),
(11, 4),
(12, 5),
(12, 6),
(13, 7),
(13, 8),
(14, 9),
(14, 10),
(15, 1),
(15, 2);



select S.FISH, S.age, S.scolarship, json_agg(C.name) as Courses, json_agg(C.teacher) as Teachers 
from Students as S
join Student_course as Sc 
on S.student_id = Sc.student_id
join Courses as C 
on Sc.course_id = C.course_id
group by S.student_id;