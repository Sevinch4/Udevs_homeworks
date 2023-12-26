create table category(
    id uuid primary key ,
    name varchar(30),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table product(
    id uuid primary key ,
    name varchar(30),
    price integer,
    category_id uuid references category(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

INSERT INTO category (id, name) VALUES
                                    ('1c0dfe8e-9d6d-42f0-8307-95b6c13c56e1', 'Electronics'),
                                    ('2a3f8d4b-1542-4d2b-9eac-8690a65cc2a2', 'Clothing'),
                                    ('3b1c5a83-2bc4-4e7e-a0a2-ff708d6d1a27', 'Home and Garden');

INSERT INTO product (id, name, price, category_id) VALUES
                                    ('4cda25d3-72a2-4c0a-a16a-09ff0c7318e8', 'Laptop', 1200, '1c0dfe8e-9d6d-42f0-8307-95b6c13c56e1'),
                                    ('5eac9e46-371f-40a8-8f1b-8d9e853a8efc', 'T-Shirt', 20, '2a3f8d4b-1542-4d2b-9eac-8690a65cc2a2'),
                                    ('6c3e3153-63c1-4f29-8ed8-4e50e7e106c5', 'Coffee Maker', 50, '3b1c5a83-2bc4-4e7e-a0a2-ff708d6d1a27'),
                                    ('84db0d84-72d0-4312-b91e-6e9e14717e70','monitor',10000000,'1c0dfe8e-9d6d-42f0-8307-95b6c13c56e1')
;
