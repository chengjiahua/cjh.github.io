CREATE DATABASE operations;

CREATE TABLE script(
    id SERIAL primary key NOT NULL,
    type varchar(32) NOT NULL,
    name varchar(64) NOT NULL,
    owner varchar(36)
);

CREATE SEQUENCE script_id_seq
START WITH 1
INCREMENT BY 1
NO MINVALUE
NO MAXVALUE
CACHE 1;

ALTER TABLE script ALTER column id SET DEFAULT nextval(script_id_seq);

