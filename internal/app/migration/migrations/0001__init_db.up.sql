CREATE SCHEMA shortener;

CREATE TABLE shortener.urls
(
    id  uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    url VARCHAR(255) NOT NULL
);
