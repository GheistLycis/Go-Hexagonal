FROM postgres:16

ARG PORT

EXPOSE ${PORT}

COPY init.sh /docker-entrypoint-initdb.d/
COPY queries/ /queries/

RUN apt-get update && apt-get install -y postgresql-client
RUN chmod +x /docker-entrypoint-initdb.d/init.sh

ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["postgres"]