create table employees(
    id uuid ,
    name varchar(30),
    phone varchar(20),
    company_id serial primary key ,
    salary int ,
    employee_type varchar(30)
);

create table company(
    id serial primary key ,
    name varchar(30),
    employees_count int
);

insert into company(id,name,employees_count) values (1,'google',12),(2,'meta',10);
