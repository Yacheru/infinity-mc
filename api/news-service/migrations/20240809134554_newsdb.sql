-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS resources (
    id serial primary key,
    url varchar default null
);

CREATE TABLE IF NOT EXISTS authors (
    id serial primary key,
    name varchar default null,
    url varchar default null,
    icon_url varchar default null
);

CREATE TABLE IF NOT EXISTS footers (
    id serial primary key,
    text text,
    icon_url varchar default null
);

CREATE TABLE IF NOT EXISTS fields (
    id serial primary key,
    name varchar,
    value varchar,
    inline boolean default false
);

CREATE TABLE IF NOT EXISTS news (
    id serial primary key,
    discord_id bigint not null,
    title text default null,
    description text default null,
    url varchar default null,
    color int default null,
    footer int default null references footers(id) on delete cascade,
    image int default null references resources(id) on delete cascade,
    thumbnail int default null references resources(id) on delete cascade,
    video int default null references resources(id) on delete cascade,
    author int default null references authors(id) on delete cascade,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

CREATE TABLE IF NOT EXISTS news_fields (
    news_id int references news(id) on delete cascade,
    field_id int references fields(id) on delete cascade
);

CREATE INDEX IF NOT EXISTS news_title_idx ON news(title);
CREATE INDEX IF NOT EXISTS news_description_idx ON news(description);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS news_fields;
DROP TABLE IF EXISTS news;
DROP TABLE IF EXISTS fields;
DROP TABLE IF EXISTS footers;
DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS resources;

DROP INDEX IF EXISTS news_title_idx;
DROP INDEX IF EXISTS news_description_idx;
-- +goose StatementEnd
