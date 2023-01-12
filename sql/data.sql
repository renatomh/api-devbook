INSERT INTO users (name, username, email, pass)
VALUES
("Jack Johnson", "jack.johson34", "jack.johson34@email.com", "$2a$10$JoVtcWwfRlQJ2WouDglZZuLcYtBUH1O83nXfcjNNePZ1T8.wpV7ou"),
("Mark Rober", "mark.rober23", "mark.rober@email.com", "$2a$10$JoVtcWwfRlQJ2WouDglZZuLcYtBUH1O83nXfcjNNePZ1T8.wpV7ou"),
("Jack Daniels", "jackdan12", "jack.dani487@email.com", "$2a$10$JoVtcWwfRlQJ2WouDglZZuLcYtBUH1O83nXfcjNNePZ1T8.wpV7ou"),
("Rup Green", "rupgr42", "rubgreen@invalid", "$2a$10$JoVtcWwfRlQJ2WouDglZZuLcYtBUH1O83nXfcjNNePZ1T8.wpV7ou"),
("Michael B.", "m.bubb2", "m.bubb2@email.com", "$2a$10$JoVtcWwfRlQJ2WouDglZZuLcYtBUH1O83nXfcjNNePZ1T8.wpV7ou");

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(2, 3),
(2, 5),
(3, 1),
(3, 2),
(3, 5);
