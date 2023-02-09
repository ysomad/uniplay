BEGIN
;

CREATE TABLE IF NOT EXISTS map (    
    id varchar(16) PRIMARY KEY NOT NULL,
    name varchar(16) UNIQUE NOT NULL,
    icon_url varchar(255) NOT NULL
);

INSERT INTO map(id, name, internal_name, icon_url)
VALUES 
    (1, 'Agency', 'cs_agency', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:Cs_agency.png'),
    (2, 'Office', 'cs_office', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:Cs_office.png'),
    (3, 'Ancient', 'de_ancient', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_ancient.png'),
    (4, 'Anubis', 'de_anubis', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_anubis.png'),
    (5, 'Cache', 'de_cache', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_cache.png'),
    (6, 'Dust II', 'de_dust2', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_dust2.png'),
    (7, 'Inferno', 'de_inferno', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_inferno.png'),
    (8, 'Mirage', 'de_mirage', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_mirage.png'),
    (9, 'Nuke', 'de_nuke', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_nuke.png'),
    (10, 'Overpass', 'de_overpass', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_overpass.png'),
    (11, 'Train', 'de_train', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_train.png'),
    (12, 'Tuscan', 'de_tuscan', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_tuscan.png'),
    (13, 'Vertigo', 'de_vertigo', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_vertigo.png'),
    (14, 'Cobblestone', 'de_cbble', 'https://developer.valvesoftware.com/wiki/List_of_CS:GO_Maps#/media/File:De_cbble.png');

CREATE TABLE IF NOT EXISTS player (
    id uuid PRIMARY KEY NOT NULL,
    steam_id numeric UNIQUE NOT NULL,
    display_name varchar(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS team (
    id smallserial PRIMARY KEY NOT NULL,
    clan_name varchar(64) UNIQUE NOT NULL,
    flag_code char(2) NOT NULL,
    captain_steam_id numeric REFERENCES player (steam_id)
);

CREATE TABLE IF NOT EXISTS team_player (
    team_id smallint NOT NULL REFERENCES team (id),
    player_steam_id numeric NOT NULL REFERENCES player (steam_id),
    is_active boolean NOT NULL DEFAULT false,
    PRIMARY KEY (team_id, player_steam_id)
);

CREATE TABLE IF NOT EXISTS match (
    id uuid PRIMARY KEY NOT NULL,
    map_name varchar(16) NOT NULL REFERENCES map (internal_name),
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

COMMIT;
