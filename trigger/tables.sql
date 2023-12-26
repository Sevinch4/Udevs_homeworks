create table employers(
    id uuid primary key ,
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

--create trigger

create function employee()
returns varchar,varchar,int,varchar
language plpgsql
as
$$
declare
 n varchar(30);
 p varchar(20);
 s int ;
 n_c varchar(30);
        begin
        select e.name into n,e.phone into p,e.salary into s,c.name into n_c from company as c join employers e on c.id = e.company_id;
    return n,p,s,n_c;