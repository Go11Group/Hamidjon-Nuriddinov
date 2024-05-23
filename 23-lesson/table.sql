-- birinchi jadval
create table Books(id serial, name varchar, author varchar, genre varchar, price real, year int);
INSERT INTO Books(name, author, genre, price, year) VALUES
('O''tkan kunlar', 'A.Qodiriy', 'Tarixiy', 35000, 1930), 
('Mehrobdan chayon', 'A.Qodiriy', 'Badiiy', 20000, 1935), 
('Diqqat', 'K.Nyuport', 'Ilmiy', 50000, 2013),
('Iymon', 'Shayx Muhammad Sodiq Muhammad Yusuf', 'Diniy', 50000,  2014), 
('1984', 'J.Oruel', 'Siyosiy', 55000, 1995), 
('Dunyoning ishlari', 'O''.Hoshimov', 'Badiiy', 25000, 2005),
('Iymon va huzun', 'S.Chamlija', 'Diniiy', 30000, 2012),
('Shaytanar', 'T.Malik', 'Detiktiv', 165000, 1985), 
('Otamdan qolgan dalalar', 'T.Murod', 'Tarixiy', 45000, 2004), 
('Ufq', 'S.Ahmad', 'Badiiy', 65000, 2009);

-- join qilish uchun ikkinchi jadval ham yaratildi
create table Kitoblar(id serial, name varchar, author varchar, genre varchar, price real, year int);
insert into books(name, author, genre, price, year) values
('Oltin zanglamas', 'P.Qodirov', 'Badiiy', 34000, 2003),
('Daftar hoshiyasidagi bitiklar', 'O''.Hoshimov', 'Badiiy', 50000, 2009),
('Jinoyat va jazo', 'F.Dostayevskey', 'Badiiy-Dramatik', 45000, 2004),
('Urush va tinchlik', 'Remark', 'Dramatik', 38000, 2014),
('Navoiy', 'Oybek', 'Tarixiy', 35000, 1965),
('Boburnoma', 'Z.M.Bobur', 'Tarixiy', 60000, 1510),
('Jimjitlik', 'S.Ahmad', 'Badiiy', 40000, 2014);
