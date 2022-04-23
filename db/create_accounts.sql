USE accmgtms;
DROP TABLE accounts;
CREATE TABLE accounts (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(128),
    DOB DATE,
    email VARCHAR(128),
    contact VARCHAR(32),
    pwd VARCHAR(128)
)