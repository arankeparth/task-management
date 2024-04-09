create database authorisation;
create database customerservice;
create database taskservice;

use authorisation;

CREATE TABLE customer_creds (
    customerid VARCHAR(255) NOT NULL PRIMARY KEY,
    username VARCHAR(255),
    password VARCHAR(255)
);


use customerservice;

CREATE TABLE customers (
    CustomerId VARCHAR(255) NOT NULL PRIMARY KEY,
    FirstName VARCHAR(255),
    LastName VARCHAR(255),
    Email VARCHAR(255) UNIQUE,
    Age INT,
    Gender VARCHAR(10)
);


use taskservice;

CREATE TABLE comments (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    content TEXT,
    ticketid INT,
    author_name VARCHAR(255),
    date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    summary VARCHAR(255),
    description TEXT,
    dueDate DATE,
    priority VARCHAR(50),
    status VARCHAR(50),
    assignee VARCHAR(100),
    dateCreated DATE,
    reporter VARCHAR(100)
);

CREATE TABLE notifications (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    type VARCHAR(255) NOT NULL,
    subject_customer_id VARCHAR(255) NOT NULL,
    customer_relation INT NOT NULL,
    subject_task_id INT,
    ts_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



