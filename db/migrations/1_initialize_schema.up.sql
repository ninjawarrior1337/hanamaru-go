CREATE TABLE IF NOT EXISTS award (
    id integer not null primary key autoincrement,
    guild_id varchar(255) not null,
    earner_id varchar(255) not null,
    award_name text not null,
    count uint not null default 1 CHECK(count >= 0)
);

CREATE UNIQUE INDEX IF NOT EXISTS award_by_guild_earner_name ON award(guild_id, earner_id, award_name);

CREATE TABLE IF NOT EXISTS tag (
    id integer not null primary key autoincrement,
    name text not null,
    guild_id varchar(255) not null,
    channel_id varchar(255) not null,
    message_id varchar(255) not null
);

CREATE UNIQUE INDEX IF NOT EXISTS tag_by_guild_tag ON tag(guild_id, name);