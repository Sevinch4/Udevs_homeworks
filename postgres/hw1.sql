create table students(
    id int, 
    first_name varchar(30),
    last_name varchar(30),
    date_of_birth date,
    math_exam_score int,
    history_exam_score int,
    literature_exam_score int,
    geography_exam_score int
);

insert into students (id, first_name, last_name, date_of_birth, math_exam_score, history_exam_score, literature_exam_score, geography_exam_score)
values
  (1, 'Nurmuhammad', 'Nuriddinov', '2004-12-12', 4, 5, 4, 3),
  (2, 'Sarvinoz', 'Mutalova', '2004-07-08', 4, 5, 5, 5),
  (3, 'Aisha', 'Bahriddinova', '2005-03-14', 5, 5, 5, 5),
  (4, 'Farangiz', 'Yusupova', '2003-09-14', 4, 5, 3, 3),
  (5, 'Qobil', 'Isakov', '2001-12-04', 3, 3, 3, 2),
  (6, 'Ilyos', 'Tojiyev', '2001-12-12', 5, 4, 3, 3),
  (7, 'Bahrom', 'Jumaboyev', '2003-01-01', 3, 4, 4, 5),
  (8, 'Toxir', 'Toxirov', '2000-05-05', 4, 4, 4, 4),
  (9, 'Islom', 'Karimov', '2003-08-07', 5, 5, 5, 4),
  (10, 'Muzaffar', 'Kamolov', '2005-09-06', 3, 5, 4, 5);

select  id, first_name, last_name, date_of_birth, math_exam_score, history_exam_score, literature_exam_score, geography_exam_score,(math_exam_score+history_exam_score + literature_exam_score + geography_exam_score) / 4.0 as average_score from students where (math_exam_score+history_exam_score + literature_exam_score + geography_exam_score)  > 3.5;

select * from students where math_exam_score < 5 and math_exam_score > 3;

select * from students where history_exam_score < 4 or history_exam_score = 4.5;