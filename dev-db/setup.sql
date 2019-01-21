DROP TABLE IF EXISTS registry;
CREATE TABLE registry (
    type_name   VARCHAR(255),
    instance_id SERIAL PRIMARY KEY,
    ip          VARCHAR(255) NOT NULL UNIQUE,
    version     VARCHAR(255),
    load        FLOAT
);