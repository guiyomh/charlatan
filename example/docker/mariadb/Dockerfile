FROM mariadb:latest

ENV MYSQL_DATABASE=fixtures \
    MYSQL_USER=fixtures_user \
    MYSQL_PASSWORD=fixtures_pass\
    MYSQL_RANDOM_ROOT_PASSWORD=true

COPY ./schema.sql docker-entrypoint-initdb.d/
