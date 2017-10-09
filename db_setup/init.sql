CREATE USER go_user WITH PASSWORD 'go_user_passwd';

CREATE DATABASE todo_test
WITH
OWNER = go_user
TABLESPACE = pg_default
CONNECTION LIMIT = -1;

\c todo_test;

CREATE TABLE todo
(
  ID          SERIAL,
  NAME        CHARACTER VARYING(255),
  DESCRIPTION VARCHAR(255),
  DUE_TO      BIGINT,
  CONSTRAINT todos_pkey PRIMARY KEY (id)
)
WITH (
OIDS = FALSE
);
ALTER TABLE todo
  OWNER TO go_user;

--Create Data

INSERT INTO todo (ID, NAME, DESCRIPTION, DUE_TO)
VALUES (1, 'GO REST API', 'Setup Go Rest API for Blog entry', 1507273200000);

INSERT INTO todo (ID, NAME, DESCRIPTION, DUE_TO)
VALUES (2, 'Create Blog Entry!', 'Create a awesome Blog entry', 1507294800000);