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

CREATE TABLE XE
(
    resource_id_td DECIMAL(18,0),
    patient_id integer,
    group_operational_id VARCHAR(18) ,
    carrier_operational_id VARCHAR(18) ,
    score_type_cde character(3) ,
    score_type_dsc VARCHAR(30) ,
    score_msr DECIMAL(5,2),
    score_msr_subgroup_cde character(5) ,
    score_subgroup_cde_dsc VARCHAR(50) ,
    disease_state_cde character(5) ,
    disease_state_txt VARCHAR(50) ,
    patient_score_eff_dte date,
    patient_score_end_dte date,
    insert_tms timestamp,
    last_update_tms timestamp,
    resource_id VARCHAR(36) NOT NULL,
    SORT KEY (patient_score_eff_dte,score_type_cde,score_msr),
    SHARD KEY (carrier_operational_id,patient_id)
);
