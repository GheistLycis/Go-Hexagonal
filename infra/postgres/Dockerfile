FROM postgres:16

COPY infra/postgres/init.sh /docker-entrypoint-initdb.d/
COPY infra/postgres/queries/ /queries/

RUN apt-get update 
RUN apt-get install -y postgresql-client
RUN chmod +x /docker-entrypoint-initdb.d/init.sh

ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["postgres"]