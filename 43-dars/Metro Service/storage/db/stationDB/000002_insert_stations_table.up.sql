INSERT INTO Stations (name) VALUES 
('Central Station'),
('North Station'),
('South Station'),
('East Station'),
('West Station'),
('Uptown Station'),
('Downtown Station'),
('Main Street Station'),
('Park Avenue Station'),
('Riverfront Station'),
('Lakeside Station'),
('Hilltop Station'),
('Valley Station'),
('Greenfield Station'),
('Pinecrest Station');


--migrate create -ext sql -dir storage/db/stationDB -seq create_stations_table
--migrate -database 'postgres://postgres:hamidjon4424@localhost:5432/metro?sslmode=disable' -path storage/db/stationDB up

