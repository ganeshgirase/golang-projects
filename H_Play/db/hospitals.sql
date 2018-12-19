/* This table contains data for hospitals */

CREATE TABLE hospitals (
  id int NOT NULL,
  name varchar(100) NOT NULL,
  address text,
  create_dt TIMESTAMP default NOW(),
  PRIMARY KEY (id)
);

