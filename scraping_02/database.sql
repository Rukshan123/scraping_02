drop database ikman_02;
create database if not exists ikman_02;

use ikman_02;

create table Adds (
    district varchar (100),
    category varchar (100),
    model varchar (100),
    price varchar (100),
    descr varchar (255),
    contact varchar (100),
    addDesc text (255)

)ENGINE=INNODB;