-- ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'Qwer1234';
-- ALTER USER 'root'@'%' IDENTIFIED BY 'Qwer1234';
-- GRANT ALL ON *.* TO 'root'@'%';
-- FLUSH PRIVILEGES;

CREATE DATABASE teszt;
use teszt;
CREATE TABLE persons ( id int PRIMARY KEY, name varchar(30) NOT NULL );
INSERT INTO persons(id,name) VALUES(1,'John Woo');
INSERT INTO persons(id,name) VALUES(2,'Peter Woo');
