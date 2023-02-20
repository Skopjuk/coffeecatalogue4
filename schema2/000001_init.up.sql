
CREATE TABLE IF NOT EXISTS users
(
    id serial primary key,
    login varchar(255) not null,
    password varchar(255) not null unique,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS coffees (
                                       id serial primary key,
                                       origin varchar(255) not null,
                                       variety varchar(255) not null,
                                       processing varchar(255) not null,
                                       descriptors varchar(255) not null,
                                       created_at timestamp not null default current_timestamp,
                                       updated_at timestamp not null default current_timestamp,
                                       roastery_id integer not null
);

CREATE TABLE IF NOT EXISTS roasteries (
                                          id serial primary key,
                                          name varchar(255) not null,
                                          country varchar(255) not null,
                                          city varchar(255) not null,
                                          address varchar(255) not null,
                                          created_at timestamp not null default current_timestamp,
                                          updated_at timestamp not null default current_timestamp
);