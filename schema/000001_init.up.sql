BEGIN;

CREATE TABLE users
(
    id serial PRIMARY KEY, -- Добавлен первичный ключ
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null 
);

CREATE TABLE todo_lists
(
    id serial PRIMARY KEY, -- Добавлен первичный ключ
    title varchar(255) not null,
    description varchar(255) 
);

CREATE TABLE users_lists
(
    id serial PRIMARY KEY, -- Добавлен первичный ключ
    user_id int references users(id) on delete cascade not null, 
    list_id int references todo_lists(id) on delete cascade not null
);

CREATE TABLE todo_items 
(
    id serial PRIMARY KEY, -- Добавлен первичный ключ
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false 
);

CREATE TABLE lists_items 
(
    id serial PRIMARY KEY, -- Добавлен первичный ключ
    item_id int references todo_items(id) on delete cascade not null, 
    list_id int references todo_lists(id) on delete cascade not null
);

COMMIT;