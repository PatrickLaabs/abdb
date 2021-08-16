CREATE DATABASE abdb CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE abdb;

CREATE TABLE tools (
  id int(11) NOT NULL AUTO_INCREMENT,
  firstname varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  lastname varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  petName varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  petSpecies varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  bloodType varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  phone int(11) DEFAULT NULL,
  plz int(11) DEFAULT NULL,
  street varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  country varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  notes text COLLATE utf8_unicode_ci,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;