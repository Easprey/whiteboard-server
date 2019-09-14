INSERT INTO users (user_id) VALUES ('test_user');

INSERT INTO fingerpaths (board, path_color, board_color, dash, blur, clear, user_id) VALUES
    ('test_board', 'FFFF0000', 'FFFFFFFF', false, false, false, 'test_user');

INSERT INTO fingerpoints (path_id, x, y) VALUES
    (LAST_INSERT_ID(), 507, 507),
    (LAST_INSERT_ID(), 508, 508),
    (LAST_INSERT_ID(), 509, 509);
    