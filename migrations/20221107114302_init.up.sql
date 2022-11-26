BEGIN;

CREATE TABLE IF NOT EXISTS team (
    name varchar(64) PRIMARY KEY NOT NULL,
    flag_code char(2) NOT NULL,
    create_time timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_time timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS player (
    steam_id bigint PRIMARY KEY NOT NULL,
    team_name varchar(64) REFERENCES team(name),
    create_time timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_time timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS team_player (
    team_name varchar(64) NOT NULL REFERENCES team (name),
    player_steam_id bigint NOT NULL REFERENCES player (steam_id),
    is_active boolean DEFAULT FALSE NOT NULL,

    PRIMARY KEY (team_name, player_steam_id)
);

CREATE TABLE IF NOT EXISTS match (
    id uuid PRIMARY KEY NOT NULL,
    map_name varchar(64) NOT NULL,
    team1_name varchar(64) NOT NULL REFERENCES team (name),
    team1_score smallint NOT NULL,
    team2_name varchar(64) NOT NULL REFERENCES team (name),
    team2_score smallint NOT NULL,
    duration interval NOT NULL,
    upload_time timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS match_player(
    match_id uuid NOT NULL REFERENCES match (id),
    player_steam_id bigint NOT NULL REFERENCES player (steam_id),
    team_name varchar(64) NOT NULL REFERENCES team (name)
);

CREATE TABLE IF NOT EXISTS metric (
    match_id uuid NOT NULL REFERENCES match (id),
    player_steam_id bigint NOT NULL REFERENCES player (steam_id),
    metric smallint NOT NULL,
    value integer NOT NULL
);

CREATE TABLE IF NOT EXISTS weapon_class (
    id smallint PRIMARY KEY NOT NULL,
    name varchar(16) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS weapon (
    id smallint PRIMARY KEY NOT NULL,
    name varchar(32) UNIQUE NOT NULL,
    class_id smallint NOT NULL REFERENCES weapon_class (id)
);

CREATE TABLE IF NOT EXISTS weapon_metric (
    match_id uuid NOT NULL REFERENCES match (id),
    player_steam_id bigint NOT NULL REFERENCES player (steam_id),
    weapon smallint NOT NULL REFERENCES weapon (id),
    metric smallint NOT NULL,
    value integer NOT NULL
);

COMMIT;