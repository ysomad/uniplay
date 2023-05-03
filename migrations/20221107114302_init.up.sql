BEGIN;

CREATE TABLE IF NOT EXISTS account (
    id uuid PRIMARY KEY NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    password varchar(4096) NOT NULL,
    is_verified boolean DEFAULT false,
    is_admin boolean DEFAULT false,
    created_at timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS map (
    name varchar(16) PRIMARY KEY NOT NULL,
    icon_url varchar(255) NOT NULL
);

INSERT INTO map(name, icon_url)
VALUES
    ('cs_agency', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:Cs_agency.png'),
    ('cs_office', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:Cs_office.png'),
    ('de_ancient', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_ancient.png'),
    ('de_anubis', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_anubis.png'),
    ('de_cache', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_cache.png'),
    ('de_dust2', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_dust2.png'),
    ('de_inferno', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_inferno.png'),
    ('de_mirage', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_mirage.png'),
    ('de_nuke', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_nuke.png'),
    ('de_overpass', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_overpass.png'),
    ('de_train', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_train.png'),
    ('de_tuscan', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_tuscan.png'),
    ('de_vertigo', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_vertigo.png'),
    ('de_cbble', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_cbble.png');

CREATE TABLE IF NOT EXISTS institution (
    id smallserial PRIMARY KEY NOT NULL,
    name varchar(255) UNIQUE NOT NULL,
    short_name varchar(255) NOT NULL,
    city varchar(128) NOT NULL,
    type smallint NOT NULL,
    logo_url varchar(2048) NOT NULL
);

CREATE TABLE IF NOT EXISTS team (
    id smallserial PRIMARY KEY NOT NULL,
    clan_name varchar(64) UNIQUE NOT NULL,
    flag_code varchar(2) NOT NULL,
    instutition_id smallint REFERENCES institution (id)
);

CREATE TABLE IF NOT EXISTS player (
    steam_id numeric PRIMARY KEY NOT NULL,
    account_id uuid UNIQUE REFERENCES account (id),
    team_id smallint REFERENCES team (id),
    display_name varchar(64),
    avatar_url varchar(2048),
    first_name varchar(32),
    last_name varchar(32)
);

CREATE TABLE IF NOT EXISTS team_player (
    team_id smallint NOT NULL REFERENCES team (id),
    player_steam_id numeric NOT NULL REFERENCES player (steam_id),
    is_captain boolean NOT NULL DEFAULT false,
    PRIMARY KEY (team_id, player_steam_id)
);

CREATE TABLE IF NOT EXISTS match (
    id uuid PRIMARY KEY NOT NULL,
    map varchar(16) NOT NULL REFERENCES map (name),
    rounds smallint NOT NULL,
    duration interval NOT NULL,
    uploaded_at timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS team_match (
    team_id smallint NOT NULL REFERENCES team (id),
    match_id uuid NOT NULL REFERENCES match (id),
    match_state smallint NOT NULL,
    score smallint NOT NULL
);

CREATE TABLE IF NOT EXISTS player_match_stat (
    player_steam_id numeric NOT NULL REFERENCES player (steam_id),
    match_id uuid NOT NULL REFERENCES match (id),
    kills smallint NOT NULL,
    hs_kills smallint NOT NULL,
    blind_kills smallint NOT NULL,
    wallbang_kills smallint NOT NULL,
    noscope_kills smallint NOT NULL,
    through_smoke_kills smallint NOT NULL,
    deaths smallint NOT NULL,
    assists smallint NOT NULL,
    flashbang_assists smallint NOT NULL,
    mvp_count smallint NOT NULL,
    damage_taken smallint NOT NULL,
    damage_dealt smallint NOT NULL,
    grenade_damage_dealt smallint NOT NULL,
    blinded_players smallint NOT NULL,
    blinded_times smallint NOT NULL,
    bombs_planted smallint NOT NULL,
    bombs_defused smallint NOT NULL
);

CREATE TABLE IF NOT EXISTS player_match (
    player_steam_id numeric NOT NULL REFERENCES player (steam_id),
    match_id uuid NOT NULL REFERENCES match (id),
    team_id smallint NOT NULL REFERENCES team (id),
    match_state smallint NOT NULL
);

CREATE TABLE IF NOT EXISTS weapon_class (
    id smallint PRIMARY KEY NOT NULL,
    class varchar(32) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS weapon (
    id smallint PRIMARY KEY NOT NULL,
    weapon varchar(32) UNIQUE NOT NULL,
    class_id smallint NOT NULL REFERENCES weapon_class (id)
);

CREATE TABLE IF NOT EXISTS player_match_weapon_stat (
    player_steam_id numeric NOT NULL REFERENCES player (steam_id),
    match_id uuid NOT NULL REFERENCES match (id),
    weapon_id smallint NOT NULL REFERENCES weapon (id),
    kills smallint NOT NULL,
    hs_kills smallint NOT NULL,
    blind_kills smallint NOT NULL,
    wallbang_kills smallint NOT NULL,
    noscope_kills smallint NOT NULL,
    through_smoke_kills smallint NOT NULL,
    deaths smallint NOT NULL,
    assists smallint NOT NULL,
    damage_taken smallint NOT NULL,
    damage_dealt smallint NOT NULL,
    shots smallint NOT NULL,
    head_hits smallint NOT NULL,
    neck_hits smallint NOT NULL,
    chest_hits smallint NOT NULL,
    stomach_hits smallint NOT NULL,
    left_arm_hits smallint NOT NULL,
    right_arm_hits smallint NOT NULL,
    left_leg_hits smallint NOT NULL,
    right_leg_hits smallint NOT NULL
);

ALTER TABLE institution ADD COLUMN ts tsvector
GENERATED ALWAYS AS
    (setweight(to_tsvector('russian', coalesce(name, '')), 'A') ||
    setweight(to_tsvector('russian', coalesce(short_name, '')), 'B')) STORED;

CREATE INDEX institution_gin_idx ON institution USING GIN (ts);

COMMIT;
