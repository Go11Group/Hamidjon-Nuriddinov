create table book
(
    id          int primary key,
    name        varchar not null,
    page        int4,
    author_name varchar,
    author_id int,
    foreign key(author_id) references author(id)
);

create table author
(
    id   int primary key,
    name varchar not null, 
);


insert into book(id, name, page, author_name, author_id) values
(1, 'O''tkan kunlar', 320, 'A.Qodiriy', 1), 
(2, 'Mehrobdan chayon', 340, 'A.Qodiriy', 1),
(3, 'Iymon', 350, 'Shayx Muhamamd Sodiq Muhammad Yusuf', 2),
(4, 'Diqqat', 280, 'K.Nyuport', 3),
(5, 'Dunyoning ishlari', 250, 'O''.Hoshimov', 4),
(6, 'Iymon va huzun', 200, 'S.Chamlija', 5),
(7, 'Uch og''ayni', 300, 'Remark', 6), 
(8, 'Shaytanat', 300, 'T.Malik', 7),
(9, 'Otamdan qolgan dalalar', 350, 'T.Murod', 8),
(10, 'Navoiy', 240, 'Oybek', 9),
(11, 'Daftar hoshiyasidagi bitiklar', 340, 'O''.Hoshimov', 4),
(12, 'Ufq', 600, 'S.Ahmad', 10),
(13, 'Jimjitlik', 350, 'S.Ahmad', 10),
(14, 'Ikki eshik orasi', 400, 'O''.Hoshimov', 4),
(15, 'Urush va tinchlik', 280, 'Remark', 6);


insert into author(id, name) values
(1, 'A.Qodiriy'),
(2, 'Shayx Muhammad Sodiq Muhammad Yusuf'),
(3, 'K.Nyuport'),
(4, 'O''.Hoshimov'),
(5, 'S.Chamlija'),
(6, 'Remark'),
(7, 'T.Malik'),
(8, 'T.Murod'),
(9, 'Oybek'),
(10, 'S.Ahmad');
