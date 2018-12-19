/* Table contsins mapping data of hospitals with their entertainment package */

CREATE TABLE hospital_category (
  hospital_id int NOT NULL,
  category varchar(20) NOT NULL,
  create_dt TIMESTAMP default NOW()
);
