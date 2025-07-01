# https://github.com/fboulnois/pg_uuidv7/blob/main/Dockerfile
FROM postgres:16

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        git \
        postgresql-server-dev-16 \
        gcc \
        make \
        libc-dev && \
        rm -rf /var/lib/apt/lists/*

RUN git clone https://github.com/fboulnois/pg_uuidv7.git /tmp/pg_uuidv7 && \
    cd /tmp/pg_uuidv7 && \
    make && \
    make install && \
    rm -rf /tmp/pg_uuidv7

RUN apt-get purge -y --auto-remove git postgresql-server-dev-16 gcc make libc-dev