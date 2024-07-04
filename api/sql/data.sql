INSERT INTO users (name, nick, email, password) 
VALUES
('User 1', 'user_1', 'user1@gmail.com', '$2a$10$P1hgrpHUSeQZnPfm5ecb.ujQGecPcaywundvcSyfZsfIkAM9yIk36'),
('User 2', 'user_2', 'user2@gmail.com', '$2a$10$P1hgrpHUSeQZnPfm5ecb.ujQGecPcaywundvcSyfZsfIkAM9yIk36'),
('User 3', 'user_3', 'user3@gmail.com', '$2a$10$P1hgrpHUSeQZnPfm5ecb.ujQGecPcaywundvcSyfZsfIkAM9yIk36');

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(1, 3);