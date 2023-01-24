BEGIN
;

CREATE TABLE IF NOT EXISTS team (
    id smallserial PRIMARY KEY NOT NULL,
    clan_name varchar(64) UNIQUE NOT NULL,
    flag_code char(2) NOT NULL
);

CREATE TABLE IF NOT EXISTS player (
    steam_id bigint PRIMARY KEY NOT NULL,
    avatar_uri varchar(2048),
    display_name varchar(24)
);

CREATE TABLE IF NOT EXISTS team_player (
    team_id smallint NOT NULL REFERENCES team (id),
    player_steam_id bigint NOT NULL REFERENCES player (steam_id),
    is_active boolean NOT NULL DEFAULT false,
    PRIMARY KEY (team_id, player_steam_id)
);

CREATE TABLE IF NOT EXISTS match (
    id uuid PRIMARY KEY NOT NULL,
    map_name varchar(64) NOT NULL,
    team1_id smallint NOT NULL REFERENCES team (id),
    team1_score smallint NOT NULL,
    team2_id smallint NOT NULL REFERENCES team (id),
    team2_score smallint NOT NULL,
    duration interval NOT NULL,
    uploaded_at timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS player_match_stat (
    player_steam_id bigint NOT NULL REFERENCES player (steam_id),
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
    player_steam_id bigint NOT NULL REFERENCES player (steam_id),
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
    player_steam_id bigint NOT NULL REFERENCES player (steam_id),
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
    chest_hits smallint NOT NULL,
    stomach_hits smallint NOT NULL,
    left_arm_hits smallint NOT NULL,
    right_arm_hits smallint NOT NULL,
    left_leg_hits smallint NOT NULL,
    right_leg_hits smallint NOT NULL
);

COMMIT;
