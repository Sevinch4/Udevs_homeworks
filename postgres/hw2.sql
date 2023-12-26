create table products (
    id serial primary key,
    name varchar(30) not null, 
    description text,
    price int ,
    created_at date,
    updated_at date,
    UNIQUE (name)
);

alter table products drop constraint products_name_key;

INSERT INTO products (name, description, price, created_at, updated_at)
VALUES
    ('Burger', 'Description for Burger', 18000, '2023-02-01', '2023-02-02'),
    ('Chocolate', 'Description for Chocolate', 7500, '2023-02-03', '2023-02-04'),
    ('Lavash', 'Description for Lavash', 26000, '2023-02-05', '2023-02-06'),
    ('Banan', 'Description for Banan', 22000, '2023-02-07', '2023-02-07'),
    ('Olma', 'Description for Olma', 500, '2023-02-09', '2023-02-09'),
    ('Kola', 'Description for Kola', 13000, '2023-02-11', '2023-02-12'),
    ('Suv', 'Description for Suv', 3000, '2023-02-13', '2023-02-13'),
    ('Snickers', 'Description for Snickers', 9500, '2023-02-15', '2023-02-16'),
    ('Pizza', 'Description for Pizza', 120000, '2023-02-17', '2023-02-18'),
    ('Yogurt', 'Description for Yogurt', 1000, '2023-02-19', '2023-02-19');

select * from products where name ilike '%a%' and price > 5000;

-- <> and != same
select * from products where (name ilike 'burger' OR description ilike '%burger%') AND created_at != updated_at;

select * from products where (name ilike 'burger' OR description ilike '%burger%') AND created_at <> updated_at;
