create table if not exists lessons (
    id integer PRIMARY KEY,
    name varchar(50) not null,
    description varchar(300) null,
    createdat date not null,
    modifiedat date,
    createdby varchar(50) not null
);

create table if not exists lessonitems (
    id integer PRIMARY KEY,
    content varchar(1500) null,
    ordinal integer not null,
    createdat date not null,
    modifiedat date null,
    createdby varchar(50) not null
);

insert into lessons (id, name, description, createdat, modifiedat, createdby)
values (1, 'Lorem ipsum', 'Foo bar', now(), NULL, 'Luke Walter');

insert into lessonitems (id, content, ordinal, createdat, modifiedat, createdby)
values (1, 'Item 1', 1, now(), NULL, 'Luke Walter');