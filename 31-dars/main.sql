create table product
(
    id       uuid,
    name     varchar,
    category varchar,
    cost     int
);


create index product_id_idx on product(id);

create index product_id_idx on product(name);

create index product_id_idx on product(category);

create index product_id_idx on product(id, name);

Explain(analyse)
Select * from product where id = 'f2aba3c9-1800-453d-bdec-e7e2cbfee23d';

Explain(analyse)
Select * from product where name = 'Prince Rylan Kshlerin';

Explain(analyse)
Select * from product where category = 'Clare';

Explain(analyse)
Select * from product where id = 'f2aba3c9-1800-453d-bdec-e7e2cbfee23d' and name = 'Lord Ernesto Rice';



-- Analiz report alohida faylda gitga yuklandi.



