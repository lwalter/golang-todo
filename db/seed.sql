create table lessons (
    id integer PRIMARY KEY,
    name varchar(50) not null,
    description varchar(300) null,
    createdat date not null,
    modifiedat date,
    createdby varchar(50) not null
); 

insert into lessons (id, name, description, createdat, modifiedat, createdby) values (1, 'Lorem ipsum', 'Foo bar', now(), NULL, 'Luke Walter');