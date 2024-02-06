CREATE TABLE users
(
    id              serial       not null unique,
    name            varchar(255) not null,
    username        varchar(255) not null unique,
    password_hash   varchar(255) not null
);

CREATE TABLE genre_lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    list_ID int references genre_lists (id) on delete cascade not null
);

CREATE TABLE films_lists
(
    id serial not null unique,
    tittle varchar(255) not null,
    description varchar(255),
    year int not null
);

CREATE TABLE genre_films
(
    id serial not null unique,
    item_id int references films_lists (id) on delete cascade not null,
    list_ID int references genre_lists (id) on delete cascade not null
);
