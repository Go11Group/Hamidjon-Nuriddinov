create table jobs(id serial primary key, name varchar);

create table workers(id serial, FISH varchar, job_id int, degre varchar, salary real, foreign key(job_id) references jobs(id));
