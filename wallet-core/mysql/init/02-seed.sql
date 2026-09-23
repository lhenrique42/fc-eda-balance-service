INSERT INTO clients (id, name, email, created_at) VALUES
    ('11111111-1111-1111-1111-111111111111', 'Alice Silva', 'alice@example.com', NOW()),
    ('22222222-2222-2222-2222-222222222222', 'Bruno Costa', 'bruno@example.com', NOW()),
    ('33333333-3333-3333-3333-333333333333', 'Carla Souza', 'carla@example.com', NOW());

INSERT INTO accounts (id, client_id, balance, created_at) VALUES
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '11111111-1111-1111-1111-111111111111', 1500.50, NOW()),
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '22222222-2222-2222-2222-222222222222', 320.00, NOW()),
    ('cccccccc-cccc-cccc-cccc-cccccccccccc', '33333333-3333-3333-3333-333333333333', 980.75, NOW());

INSERT INTO transactions (id, account_id_from, account_id_to, amount, created_at) VALUES
    ('dddddddd-dddd-dddd-dddd-dddddddddddd', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 100.00, NOW()),
    ('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'cccccccc-cccc-cccc-cccc-cccccccccccc', 50.25, NOW()),
    ('ffffffff-ffff-ffff-ffff-ffffffffffff', 'cccccccc-cccc-cccc-cccc-cccccccccccc', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 200.00, NOW());
