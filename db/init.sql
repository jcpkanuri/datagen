use information_schema;

create user 'developer'@'%' identified by 'tester123';
create database company;
grant all on company.* to 'developer'@'%';
use company;
create table employee( id int not null auto_increment, ename varchar(100), primary key(id));
insert into employee (ename) values('john');
insert into employee (ename) values('joe');
insert into employee (ename) values('jessy');


create table department(
    id int not null auto_increment,
    name varchar(10), primary key(id)
);

insert into department (name) values('presales');
insert into department (name) values('design');
insert into department (name) values('sales');
insert into department (name) values('account');
insert into department (name) values('billing');

alter table employee add column dept varchar(10);

create table numeric_types (
    id int not null auto_increment,
    mytinyint tinyint,
    mysmallint smallint,
    mymediumint mediumint,
    myint int,
    mybigint bigint,
    myfloat float,
    mydouble double,
    mydecimal decimal,
     primary key(id)    
);

create table time_types (    
    id int not null auto_increment,
    mydate date,
    mytime time,
    mytstamp timestamp,
    myyear year,
    primary key(id) 
);

create table char_types (
    id int not null auto_increment,
    mychar char(5),
    mybinary binary(64),
    myvarchar varchar(200),
    myvarbinary varbinary(200),
    mytext longtext,
    primary key(id) 
);

create ROWSTORE table geo_types (
    id int not null auto_increment,
    gpoint geographypoint,
    geo geography,
    primary key(id) 
);

create ROWSTORE table seq_test (
    id int not null,
    seq1 int ,
    seq2 int,
    primary key(id) 
);
