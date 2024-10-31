create table if not exists brokerage_accounts (
    id bigserial primary key,
    name text not null unique,
    balance decimal not null
);