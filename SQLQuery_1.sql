CREATE DATABASE Manager;

CREATE TABLE Account (
    id INT PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE Informations (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    fullname VARCHAR(255),
    gender VARCHAR(255),
    email VARCHAR(255),
    CreateAt DATETIME,
    UpdateAt DATETIME,
    FOREIGN KEY (id) REFERENCES Account(id)
);

CREATE TABLE Role (
    id INT PRIMARY KEY,
    fullname VARCHAR(255),
    roleID INT,
    FOREIGN KEY (id) REFERENCES Account(id)
);

CREATE TABLE Premissions (
    id INT PRIMARY KEY,
    roleID INT,
    name_role VARCHAR(255),
    premissions VARCHAR(255),
    FOREIGN KEY (id) REFERENCES Account(id)
);

-- Tạo 2 đối tượng cho bảng Account
INSERT INTO Account (id, email, password) VALUES (1, 'langbach364@gmail.com', 'password1');
INSERT INTO Account (id, email, password) VALUES (2, 'langbach363@gmail.com', 'password2');
INSERT INTO Account (id, email, password) VALUES (3, 'langbach362@gmail.com', 'password3');


-- Tạo 2 đối tượng cho bảng Informations
INSERT INTO Informations (id, name, fullname, gender, email, CreateAt, UpdateAt) VALUES (1, 'Name1', 'Fullname1', 'Gender1', 'email1@example.com', NOW(), NOW());
INSERT INTO Informations (id, name, fullname, gender, email, CreateAt, UpdateAt) VALUES (2, 'Name2', 'Fullname2', 'Gender2', 'email2@example.com', NOW(), NOW());
INSERT INTO Informations (id, name, fullname, gender, email, CreateAt, UpdateAt) VALUES (3, 'Name2', 'Fullname2', 'Gender2', 'email2@example.com', NOW(), NOW());

-- Tạo 2 đối tượng cho bảng Role
INSERT INTO Role (id, fullname, roleID) VALUES (1, 'Fullname1', 1);
INSERT INTO Role (id, fullname, roleID) VALUES (2, 'Fullname2', 2);
INSERT INTO Role (id, fullname, roleID) VALUES (3, 'Fullname2', 2);

-- Tạo 2 đối tượng cho bảng Premissions
INSERT INTO Premissions (id, roleID, name_role, premissions) VALUES (1, 1, 'Role1', 'Premission1');
INSERT INTO Premissions (id, roleID, name_role, premissions) VALUES (2, 2, 'Role2', 'Premission2');
INSERT INTO Premissions (id, roleID, name_role, premissions) VALUES (3, 2, 'Role2', 'Premission2');

SELECT  * from Account
DROP TABLE IF EXISTS Premissions;
DROP TABLE IF EXISTS Role;
DROP TABLE IF EXISTS Informations;
DROP TABLE IF EXISTS Account;

SELECT password FROM Account WHERE email = "langbach364@gmail.com"
delete from Account WHERE id = 1