create or replace function trigger_func()
       returns trigger as
       $$
begin
        if not exists (select 1 from company where id = new.company_id) then
           insert into company (id,name,employees_count)
           values(new.company_id,
            substr(md5(random() :: text),1,30),
            floor(random() * 100 +1)
           );
end if;
return new;
end;
           $$ language plpgsql;


create trigger check_company_id
    before insert or update on employees
                         for each row
                         execute function trigger_func();


insert into employees (id,name,phone,company_id,salary,employee_type) values('fab8a9eb-01fd-4672-9430-2a2f77a2136a','sarvar','908887656',3,300000,'middle');

--function
create function get_employee_info()
    returns table(name varchar,phone varchar,salary int,company_name varchar)
    language plpgsql as
    $$
begin
        return query select e.name,e.phone ,e.salary,c.name
            from company as c
            inner join employees as e on c.id = e.company_id;

end;
    $$