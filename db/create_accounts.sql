CREATE TABLE accounts (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(128),
    DOB TIMESTAMP,
    email VARCHAR(128),
    contact VARCHAR(32),
    pwd PASSWORD(32)
)