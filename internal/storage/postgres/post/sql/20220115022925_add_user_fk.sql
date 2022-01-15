-- +goose Up
-- +goose StatementBegin
alter table "post"
    add column owner bigint,
    add constraint fk_owner foreign key (owner) references "user" (ident)
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table "post"
    drop constraint fk_owner,
    drop column owner
;
-- +goose StatementEnd
