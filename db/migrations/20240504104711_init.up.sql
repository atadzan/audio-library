CREATE TABLE IF NOT EXISTS users
(
    id serial primary key not null,
    fullname varchar(255) not null,
    login varchar(255) not null unique,
    password_hash varchar(255) not null,
    created_at timestamp(0) not null
);

CREATE TABLE IF NOT EXISTS tracks
(
    id bigserial primary key not null,
    title varchar(255) not null,
    artist varchar(255) not null,
    genre varchar not null,
    path varchar(255) not null unique ,
    uploader_id int references users(id) not null,
    created_at timestamp(0) not null
);

CREATE TABLE IF NOT EXISTS liked_tracks
(
    track_id bigint references tracks(id) not null,
    user_id int references users(id) not null,
    created_at timestamp(0) not null,
    constraint user_track_uq unique (track_id, user_id)
);
