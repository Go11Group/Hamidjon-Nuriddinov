INSERT INTO Users (name, phone, age) VALUES 
('Alice Smith', '123-456-7890', 30),
('Bob Johnson', '234-567-8901', 25),
('Charlie Brown', '345-678-9012', 35),
('David Williams', '456-789-0123', 28),
('Eve Davis', '567-890-1234', 32),
('Frank Miller', '678-901-2345', 40),
('Grace Wilson', '789-012-3456', 22),
('Hannah Moore', '890-123-4567', 27),
('Ian Taylor', '901-234-5678', 33),
('Jack White', '012-345-6789', 29),
('Karen Harris', '123-987-6543', 31),
('Larry Clark', '234-876-5432', 26),
('Monica Lewis', '345-765-4321', 36),
('Nate Hall', '456-654-3210', 34),
('Olivia Young', '567-543-2109', 23),
('Paul King', '678-432-1098', 38),
('Quinn Scott', '789-321-0987', 24),
('Rachel Adams', '890-210-9876', 29),
('Sam Baker', '901-109-8765', 37),
('Tina Perez', '012-098-7654', 21);

--migrate create -ext sql -dir storage/db -seq create_users_table
--migrate -database 'postgres://postgres:hamidjon4424@localhost:5432/users?sslmode=disable' -path storage/db up
