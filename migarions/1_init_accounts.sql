create table if not exists accounts (
    id bigserial primary key,
    name text not null unique,
    balance bigint not null
);