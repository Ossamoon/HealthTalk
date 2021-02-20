USE go_api_mysql;

CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users VALUES (1, 'sampleman', 'samplepass');
INSERT INTO users VALUES (2, 'testman', 'testpass');