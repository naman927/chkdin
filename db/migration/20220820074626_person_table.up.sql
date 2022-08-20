create table if not EXISTS persons(
  id int not null key auto_increment,
  firstname varchar(50),
  lastname varchar(50),
  address varchar(50),
  contactno varchar(50),
  email VARCHAR(50)
);