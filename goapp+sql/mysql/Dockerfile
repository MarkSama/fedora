# Alap image MySQL-hez

FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=Qwer1234

ADD ./init.sql /docker-entrypoint-initdb.d/

#EXPOSE 3306

CMD ["mysqld"]