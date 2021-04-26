DROP DATABASE student_dosen_pa;
CREATE DATABASE student_dosen_pa;
USE student_dosen_pa;
DROP TABLE student;
DROP TABLE dosen_pa;

CREATE TABLE dosen_pa (
      dosen_id     BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
      name         VARCHAR(100) NOT NULL,
      identifier   CHAR(10) NOT NULL UNIQUE ,
      email        VARCHAR(100) NOT NULL UNIQUE,
      age          INT NOT NULL
);

CREATE TABLE student (
       student_id   BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
       name         VARCHAR(100) NOT NULL,
       identifier   CHAR(10) NOT NULL UNIQUE ,
       email        VARCHAR(100) NOT NULL UNIQUE,
       age          INT NOT NULL,
       dosen_pa_id  BIGINT,
       constraint foreign key (dosen_pa_id)  REFERENCES dosen_pa(dosen_id) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE = InnoDB;

INSERT INTO dosen_pa(name, identifier, email, age) VALUES
('sammidev1', '1111111111', 'sammidev1', 26),
('sammidev2', '1111111112', 'sammidev2', 25),
('sammidev3', '1111111113', 'sammidev3', 27),
('sammidev4', '1111111114', 'sammidev4', 28),
('sammidev5', '1111111115', 'sammidev5', 29);


INSERT INTO student(name, identifier, email, age, dosen_pa_id) VALUES
('sam1', '1111111111', 'sam1', 19, 1),
('sam2', '1111111112', 'sam2', 19, 1),
('sam3', '1111111113', 'sam3', 19, 2),
('sam4', '1111111114', 'sam4', 19, 2),
('sam5', '1111111115', 'sam5', 19, 3),
('sam6', '1111111116', 'sam6', 19, 3),
('sam7', '1111111117', 'sam7', 19, 3),
('sam8', '1111111118', 'sam8', 19, 3),
('sam9', '1111111119', 'sam9', 19, 3),
('sam10','1111111110', 'sam10', 19, 4),
('sam11','1111111121', 'sam11', 19, 5);

SELECT p.student_id, p.name, p.identifier, p.email, p.age, p.dosen_pa_id, dp.dosen_id, dp.name, dp.identifier, dp.email, dp.age FROM student p JOIN dosen_pa dp on p.dosen_pa_id = dp.dosen_id;