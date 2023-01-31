CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE coffee
(
    id serial not null unique,
    origin varchar(255) not null,
    variety varchar(255) not null,
    processing varchar(255) not null,
    descriptors varchar(255) not null
)