USE go_api_mysql;

CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE talks (
    id INT NOT NULL AUTO_INCREMENT,
    uid_from INT NOT NULL,
    uid_to INT NOT NULL,
    content VARCHAR(1024) NOT NULL, 
    PRIMARY KEY (id)
);

INSERT INTO users (id, name, password) VALUES (1, 'sampleman', 'samplepass'), (2, 'testman', 'testpass');
INSERT INTO talks (uid_from, uid_to, content) VALUES (1, 2, 'hello!!'), (2, 1, 'goodmorning!');