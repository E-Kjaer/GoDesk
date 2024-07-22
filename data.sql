BEGIN;
CREATE TABLE workcards (
    id SERIAL PRIMARY KEY,
    status VARCHAR NOT NULL
);

CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    value VARCHAR NOT NULL
);

CREATE TABLE workcardtags (
    workcard INT references workcards(id),
    tag INT references tags(id),
    PRIMARY KEY (workcard, tag)
);

INSERT INTO workcards (status) VALUES ('Ongoing');
INSERT INTO workcards (status) VALUES ('Ongoing');
INSERT INTO workcards (status) VALUES ('Ongoing');

INSERT INTO tags (value) VALUES ('New Bike');
INSERT INTO tags (value) VALUES ('Service');

INSERT INTO workcardtags (workcard, tag) VALUES (1, 1);
INSERT INTO workcardtags (workcard, tag) VALUES (1, 2);

SELECT workcards.id, status, value FROM workcards
    JOIN workcardtags on workcardtags.workcard = workcards.id
    JOIN tags t ON workcardtags.tag = t.id;

ROLLBACK;