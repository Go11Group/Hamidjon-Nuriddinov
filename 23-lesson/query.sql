-- Update - jadval ichidagi ma;lumotlarni qiymatlarini o'zgartirish uchun ishlatiladi

Update from Books
Set price = 60000 
where id = 3; 


-- Delete - jadvaldagi ma'lumotlarni qiymatlarini o'chirish uchun ishlatiladi. 

Delete from Books;
-- Barcha ma'lumotlarni o'chiradi

Delete from Books 
where price < 40000;
-- Narxi 40000 dan kichik bo'lgan row dagi ma'lumotlarni o;chiradi

-- Group BY - ma'lumotlarni qiymatlari bo'yicha guruhlash, ya'ni berilgan column qiymatlari bo'yicha bir xillarini guruhlaydi
-- Asosan agregat fubksiyalar bilan ishlatiladi

Select genre from Books
Group By genre;

-- Order by - ma'lumotlarning qiymatlarini berilgan column bo'yicha tartiblab chiqaradi

Select * from Books
Order BY price
-- o'sish tartibida

Select * from Books
Order By price desc 
--teskari tartibda

-- Join - jadvallarni qator ko'rinishida qo'shadi
-- join turlari:

-- Inner join - ikkita jadvaldagi umumiy qiymatlarni oladi

Select * from Books as B
Inner Join Kitoblar as K
ON B.genre = K.genre

-- Left join - birinchi jadvaldagi barcha ma'lumotlar, 
-- ikkinchi jadvaldan esa faqat berilgan columnga mos birinchi jadvalda bo'lgan elementlar chiqadi

Select * from Books as B
Left Join Kitoblar as K
On B.author = K.author

-- Right join - ikkinchi jadvaldagi barcha ma'lumotlar, 
-- birinchi jadvaldan esa faqat berilgan columnga mos ikkinchi jadvalda bo'lgan elementlar chiqadi

Select * from Books as B
Right Join Kitoblar as K
On B.author = K.author

-- Full join - ikki jadvalni ham qo'shib barcha ma'lumotlarni chiqaradi

Select * from Books as B
Full Join Kitoblar as K
On B.author = K.author