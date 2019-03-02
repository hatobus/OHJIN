from mysql:5.7

COPY DB/init/* /docker-entrypoint-initdb.d/

CMD ["mysqld"]
