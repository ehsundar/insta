create table if not exists "post"
(
    ident      bigserial primary key,
    caption    varchar(4000) not null,
    images     varchar(120)[],
    created_at timestamp default now()
);

create table if not exists "user"
(
    ident      bigserial primary key,
    username   varchar(128) unique not null,
    password   varchar(512)        not null,
    created_at timestamp default now()
);
