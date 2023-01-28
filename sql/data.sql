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

INSERT INTO posts (title, content, author_id)
VALUES
("Jack Johnon's Post", "This is Jack Johnon's Post! Cool beans, bro!", 1),
("Mark Rober's Post", "This is Mark Rober's Post! Off the charts, man!", 2),
("Jack Daniels's Post", "This is Jack Daniels's Post! Keep walking!", 3),
("Rup Green's Post", "This is Rup Green's Post! To infinity, and beyond!", 4),
("Michael B.'s Post", "This is Michael B.'s Post! Ay, mate!", 5),
("Jack Johnon's Post", "This is another Jack Johnon's Post! Cool beans, bro!", 1),
("Mark Rober's Post", "This is another Mark Rober's Post! Off the charts, man!", 2),
("Jack Daniels's Post", "This is another Jack Daniels's Post! Keep walking!", 3),
("Rup Green's Post", "This is another Rup Green's Post! To infinity, and beyond!", 4),
("Michael B.'s Post", "This is another Michael B.'s Post! Ay, mate!", 5),
("Jack Johnon's Post", "This is yet another Jack Johnon's Post! Cool beans, bro!", 1),
("Mark Rober's Post", "This is yet another Mark Rober's Post! Off the charts, man!", 2),
("Jack Daniels's Post", "This is yet another Jack Daniels's Post! Keep walking!", 3),
("Rup Green's Post", "This is yet another Rup Green's Post! To infinity, and beyond!", 4),
("Michael B.'s Post", "This is yet another Michael B.'s Post! Ay, mate!", 5);
