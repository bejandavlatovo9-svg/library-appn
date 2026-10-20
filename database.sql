CREATE  TABLE noteooks(
    id SERIAL PRIMARY KEY ,
    title VARCHAR(100) NOT NULL,
    author VARCHAR(100) NOT NULL ,
    price NUMERIC(10,2) NOT NULL ,
    pages INTEGER
);

INSERT INTO noteooks(title, author, price, pages)
VALUES
    ('Clean code', 'Bezhan Davlat',10,291),
    ('dzjhzjlhf', 'Nurulloev Nurullo',100,200);

SELECT  *FROM noteooks;
