BEGIN;

INSERT INTO weapon_class(id, name)
VALUES 
    (1, 'pistol'),
    (2, 'smg'),
    (3, 'heavy'),
    (4, 'rifle'),
    (5, 'equipment'),
    (6, 'grenade');

INSERT INTO weapon(id, name, class_id)
VALUES
    (1, 'P2000', 1),
    (2, 'Glock', 1),
    (3, 'P250', 1),
    (4, 'Desert Deagle', 1),
    (5, 'Five Seven', 1),
    (6, 'Dual Berettas', 1),
    (7, 'Tec-9', 1),
    (8, 'CZ75-Auto', 1),
    (9, 'USP-S', 1),
    (10, 'R8 Revolver', 1),

    (101, 'MP7', 2),
    (102, 'MP9', 2),
    (103, 'PP-Bizon', 2),
    (104, 'MAC-10', 2),
    (105, 'UMP-45', 2),
    (106, 'P90', 2),
    (107, 'MP5-SD', 2),

    (201, 'Sawed-Off', 3),
    (202, 'Nova', 3),
    (203, 'MAG-7', 3),
    (204, 'XM1014', 3),
    (205, 'M249', 3),
    (206, 'Negev', 3),

    (301, 'Galil AR', 4),
    (302, 'FAMAS', 4),
    (303, 'AK-47', 4),
    (304, 'M4A4', 4),
    (305, 'M4A1-S', 4),
    (306, 'SSG 08', 4),
    (307, 'SG 553', 4),
    (308, 'AUG', 4),
    (309, 'AWP', 4),
    (310, 'SCAR-20', 4),
    (311, 'G3SG1', 4),

    (401, 'Zeus x27', 5),
    (404, 'C4', 5),
    (405, 'Knife', 5),
    (407, 'World', 5),

    (501, 'Decoy Grenade', 6),
    (502, 'Molotov', 6),
    (503, 'Incendiary Grenade', 6),
    (504, 'Flashbang', 6),
    (505, 'Smoke Grenade', 6),
    (506, 'HE Grenade', 6);

COMMIT;

