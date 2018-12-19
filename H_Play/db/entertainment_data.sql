/* Table contains information related to entertainment */

CREATE TABLE  entertainment_data (
  category varchar(20) NOT NULL,
  title varchar(50) NOT NULL,
  provider_email VARCHAR(40),
  description jsonb,
  create_dt TIMESTAMP default NOW(),
  PRIMARY KEY (category,title)
);

