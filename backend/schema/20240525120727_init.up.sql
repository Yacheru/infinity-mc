CREATE TABLE active_payments (
    id serial not null unique,
    payment_id varchar(255) not null unique
);

CREATE TABLE payments_history (
    id serial not null unique,
    payment_id varchar(255) not null unique,
    price varchar(255) not null,
    donat_type varchar(255) not null
);

CREATE TABLE users_payments (
    id serial not null unique,
    nickname varchar(255) not null,
    payment_id int references payments_history (id) on delete cascade not null
)
