CREATE TABLE users(
    id serial PRIMARY KEY NOT NULL,
    name varchar,
    email varchar,
    password varchar
);

CREATE TABLE debts(
    from_id serial REFERENCES users(id),
    to_id serial REFERENCES users(id),
    sums integer
);

CREATE TABLE history_checks(
    id serial PRIMARY KEY NOT NULL ,
    author_id serial REFERENCES users(id)
);

CREATE TABLE products(
    check_id serial REFERENCES history_checks(id),
    name varchar,
    quantity integer,
    price integer,
    users_ids varchar
);