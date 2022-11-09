CREATE TABLE IF NOT EXISTS player (
    steam_id bigint PRIMARY KEY NOT NULL,
    is_created_manually boolean DEFAULT FALSE NOT NULL,
    create_time timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_time timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS team (
    name varchar(16) PRIMARY KEY NOT NULL,
    flag_code char(2) NOT NULL,
    is_created_manually boolean DEFAULT FALSE NOT NULL,
    create_time timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_time timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE TABLE IF NOT EXISTS team_player (
    team_name varchar(16) PRIMARY KEY NOT NULL REFERENCES team (name),
    player_steam_id bigint PRIMARY KEY NOT NULL REFERENCES player (steam_id),
    is_active boolean DEFAULT FALSE NOT NULL
);

CREATE TABLE IF NOT EXISTS match (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
    map_name varchar(64) NOT NULL,
    team1_name varchar(16) NOT NULL REFERENCES team (name),
    team1_score smallint NOT NULL,
    team2_name varchar(16) NOT NULL REFERENCES team (name),
    team2_score smallint NOT NULL,
    duration interval NOT NULL,
    upload_time timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS metric (
    match_id uuid NOT NULL REFERENCES match (id),
    player_steam_id bigint NOT NULL REFERENCES player (steam_id),
    metric smallint NOT NULL,
    value integer NOT NULL
);

CREATE TABLE IF NOT EXISTS weapon_metric (
    match_id bigint NOT NULL REFERENCES match (id),
    player_steam_id bigint NOT NULL REFERENCES player (steam_id),
    weapon_name varchar(64) NOT NULL,
    weapon_class varchar(64) NOT NULL,
    metric smallint NOT NULL,
    value integer NOT NULL
);