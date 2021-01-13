CREATE TABLE users
(
    id       bigserial      NOT NULL primary key,
    username varchar UNIQUE NOT NULL,
    password varchar        NOT NULL
);

CREATE TABLE todos
(
    id               bigserial NOT NULL primary key,
    user_id          bigserial NOT NULL,
    title            varchar   NOT NULL,
    description      varchar   NOT NULL,
    time_to_complete timestamp not null,
    constraint "user_id_fk" foreign key ("user_id") references "users" ("id")
);