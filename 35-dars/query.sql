Create table Problem(id uuid default gen_random_uuid(), status bool, name varchar, difficulty varchar);

Create table Users(id uuid default gen_random_uuid(), name varchar, programming_language varchar);

Create table Solved_problem(id uuid default gen_random_uuid(), user_id uuid, problem_id uuid);
