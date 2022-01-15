-- +goose Up
-- +goose StatementBegin
create table post
(
    ident      bigserial primary key,
    caption    varchar(4000) not null,
    images     varchar(120)[],
    created_at timestamp default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "post";
-- +goose StatementEnd
