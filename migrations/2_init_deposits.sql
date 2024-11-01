create table if not exists deposits (
    id bigserial primary key,
    name text not null unique,
    balance decimal not null,
    rate decimal not null,
    month_period bigint not null,
    created_at timestamp with time zone not null,
    closed_at timestamp with time zone null
)