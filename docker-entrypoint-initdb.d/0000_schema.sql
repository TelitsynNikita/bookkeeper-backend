create table users
(
    id        serial unique,
    full_name varchar(255) not null,
    email     varchar(255) not null unique,
    password  varchar(255) not null,
    role      varchar(255) default 'USER'
);

create table requests
(
    id serial unique,
    purpose varchar(255) not null,
    description varchar(255) not null,
    amount integer not null,
    status varchar(255) default 'NEW' not null,
    user_id integer not null references users (id) on delete cascade,
    bookkeeper_id integer references users (id) on delete cascade default 0,
    created_at    timestamp    default now()
);