CREATE DATABASE Manager;

USE Manager

CREATE TABLE Account (
    id INT PRIMARY KEY,
    email VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE Informations (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    fullname VARCHAR(255),
    gender VARCHAR(255),
    email VARCHAR(255),
    CreateAt DATETIME,
    UpdateAt DATETIME,
    FOREIGN KEY (id) REFERENCES Account(id) ON DELETE CASCADE
);

CREATE TABLE RoleDefinitions (
    roleID INT PRIMARY KEY, 
    name_role VARCHAR(255)
);


CREATE TABLE Role (
    id INT PRIMARY KEY,
    fullname VARCHAR(255),
    roleID INT,
    FOREIGN KEY (id) REFERENCES Account(id) ON DELETE CASCADE,
    FOREIGN KEY (roleID) REFERENCES RoleDefinitions(roleID)
);


CREATE TABLE Token (
    id INT,
    tokenID VARCHAR(255) PRIMARY KEY,
    code VARCHAR(255),
    FOREIGN KEY (id) REFERENCES Role(id) ON DELETE CASCADE
);

CREATE TABLE Premissions (
    id INT PRIMARY KEY,
    roleID INT,
    premissions VARCHAR(255),
    FOREIGN KEY (id) REFERENCES Account(id) ON DELETE CASCADE,
    FOREIGN KEY (roleID) REFERENCES RoleDefinitions(roleID)
);


-- Chèn vào bảng Account
INSERT INTO Account(id, email, password)
VALUES (1, 'example@email.com', 'password123');

-- Chèn vào bảng Informations
INSERT INTO Informations(id, name, fullname, gender, email, CreateAt, UpdateAt)
VALUES (1, 'Tên', 'Họ và tên', 'Nam', 'example@email.com', NOW(), NOW());

-- Chèn vào bảng RoleDefinitions
INSERT INTO RoleDefinitions(roleID, name_role)
VALUES (1, 'Admin');

-- Chèn vào bảng Role
INSERT INTO Role(id, fullname, roleID)
VALUES (1, 'Họ và tên', 1);

-- Chèn vào bảng Token
INSERT INTO Token(id, tokenID, code)
VALUES (1, 'token123', 'code123');

-- Chèn vào bảng Premissions
INSERT INTO Premissions(id, roleID, premissions)
VALUES (1, 1, 'full_access');


-- Chèn vào bảng Account
INSERT INTO Account(id, email, password)
VALUES (1, 'example@email.com', 'password123');

-- Chèn vào bảng Informations
INSERT INTO Informations(id, name, fullname, gender, email, CreateAt, UpdateAt)
VALUES (1, 'Tên', 'Họ và tên', 'Nam', 'example@email.com', NOW(), NOW());

-- Chèn vào bảng RoleDefinitions
INSERT INTO RoleDefinitions(roleID, name_role)
VALUES (1, 'Admin');

-- Chèn vào bảng Role
INSERT INTO Role(id, fullname, roleID)
VALUES (1, 'Họ và tên', 1);

-- Chèn vào bảng Token
INSERT INTO Token(id, tokenID, code)
VALUES (1, 'token123', 'code123');

-- Chèn vào bảng Premissions
INSERT INTO Premissions(id, roleID, premissions)
VALUES (1, 1, 'full_access');



DELETE FROM Account WHERE id = 1
SELECT * from Account
SELECT * from Premissions
SELECT id FROM Token WHERE code = code1


drop DATABASE Manager;