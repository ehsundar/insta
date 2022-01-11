create table "user"
(
    ident      BIGSERIAL PRIMARY KEY,
    username   varchar(128) unique not null,
    password   varchar(512)        not null,
    created_at timestamp default now()
);
