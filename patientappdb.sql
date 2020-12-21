CREATE DATABASE patientappdb;
USE patientappdb;

CREATE TABLE users(
	username varchar(255),
    password varchar(255)
);

CREATE TABLE patients(
	id varchar(255),
    name varchar(255),
    dob varchar(355),
    phone varchar(255),
    email varchar(255),
    address varchar(255),
    time timestamp,
    PRIMARY KEY(id)
);

select * from patientappdb.users;

select * from patientappdb.patients;