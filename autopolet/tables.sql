create table ticket(
    id uuid primary key ,
    departure varchar(40),
    destination varchar(40),
    date timestamp
);

create table users(
    id uuid primary key ,
    first_name varchar(40),
    last_name varchar(40),
    email varchar(40),
    phone varchar(13),
    ticket_id uuid references ticket(id)
);

insert into ticket values('c17efe64-b0eb-4d5c-badc-f946151d5858','sidney','vietnam',current_timestamp),
                         ('0545f77c-ee9d-4990-ae83-31e7132e9d4c','saudiya','parij',current_timestamp),
                         ('f7181123-5be7-4819-ac93-024c5aef0b2b','malayziya','antalya',current_timestamp),
                         ('07981ddf-e2b9-4fb6-87b5-0c8db139c297','parij','italy',current_timestamp);

insert into users values ('968d78ef-76af-4dbb-a1bd-ff658b61b3fd','Qobil','Toxirov','qobil@gmail.com','99-860-90-09','c17efe64-b0eb-4d5c-badc-f946151d5858'),
                       ('f4275bf2-a3f5-41f1-b555-066fb53660a6','Nurmuhammad','Nuriddinov','nur@gmail.com','33-444-88-77','07981ddf-e2b9-4fb6-87b5-0c8db139c297'),
                       ('39443a1a-ab6e-44b9-bc46-e5a7c4673911','Farangiz','Mirjalolava','Mfarangiz@gmail.com','90-888-67-56','f7181123-5be7-4819-ac93-024c5aef0b2b'),
                       ('e640001f-c0df-4c4d-b001-14735cf8191c', 'John', 'Doe', 'john.doe@example.com', '123-456-7890', 'c17efe64-b0eb-4d5c-badc-f946151d5858'),
                       ('f640001f-c0df-4c4d-b001-14735cf8191d', 'Jane', 'Doe', 'jane.doe@example.com', '987-654-3210', 'f7181123-5be7-4819-ac93-024c5aef0b2b'),
                       ('c2ca7c4c-0969-45d8-a892-dd3c986d6f7b', 'Alice', 'Smith', 'alice.smith@example.com', '111-222-3333','0545f77c-ee9d-4990-ae83-31e7132e9d4c'),
                       ('c9c0af87-c767-47c1-8dcd-bfcd523bcd36', 'Abboshon', 'Jorayev', 'abbos@example.com', '33-124-11-11', 'f7181123-5be7-4819-ac93-024c5aef0b2b'),
                       ('00f8d8cc-ca60-4b3b-9396-e0b63f3096cf', 'Sarvar', 'Bazarov', 'sarvar.baz@example.com', '99-001-11-90','0545f77c-ee9d-4990-ae83-31e7132e9d4c');