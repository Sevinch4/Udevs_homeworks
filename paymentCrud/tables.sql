create table payments(
    id uuid primary key ,
    payment_type varchar(30),
    date timestamp,
    amount int
);

INSERT INTO payments (id, payment_type, date, amount) VALUES
                                                          ('b1c22a3e-4e1a-4e86-8a7b-85a09a37aefd', 'Payme', '2023-01-01 10:30:00', 500),
                                                          ('f2b3e6e1-3e7d-40a0-93b7-2a3bcf7f678a', 'Cash', '2023-01-02 15:45:00', 300),
                                                          ('d7c4af3d-1f6c-4c80-825d-7a695a3d6d74', 'PayPal', '2023-01-03 08:00:00', 700),
                                                          ('e8a150de-9d58-4d37-a186-32b245a43b11', 'Payme', '2023-01-04 12:15:00', 1000),
                                                          ('a9b6e17d-4c47-437d-8c82-6436a78d456a', 'Cash', '2023-01-05 18:20:00', 250),
                                                          ('c6e51d2b-29e0-4f46-9db8-13c6412335ef', 'PayPal', '2023-01-06 09:45:00', 450),
                                                          ('fdaa3418-929b-4d05-aa66-d9ea5b22f543', 'Payme', '2023-01-07 14:30:00', 800),
                                                          ('a6e05fb8-4e7b-4b4e-9c6c-7865ff979536', 'Cash', '2023-01-08 16:00:00', 600),
                                                          ('cfda86b5-7dd5-46c7-bf14-13a824e2cc45', 'PayPal', '2023-01-09 11:10:00', 350),
                                                          ('b6e418bd-0a34-4fcf-9b68-2796da85a1bd', 'Payme', '2023-01-10 20:45:00', 1200);