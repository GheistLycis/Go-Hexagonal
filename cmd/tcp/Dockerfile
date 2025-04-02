FROM golang:1.24 AS base

WORKDIR /app

COPY .env .
COPY scripts/build.sh .
COPY src/ src/
COPY cmd/ cmd/
COPY infra/ infra/
COPY main.go .

RUN apt-get -y update
RUN go mod init Go-Hexagonal
RUN go mod tidy
RUN HOST_IP=$(getent hosts host.docker.internal | awk '{ print $1 }' || echo "127.0.0.1") && \
    echo "Detected Host IP: $HOST_IP" && \
    export ALLOWED_IPS="$ALLOWED_IPS,$HOST_IP" && \
    echo "ALLOWED_IPS=$ALLOWED_IPS" >> /etc/environment
RUN chmod +x ./build.sh
RUN ./build.sh tcp main

ENTRYPOINT ["bin/main"]


# TODO: fix "exec ./main: no such file or directory"
# FROM scratch

# WORKDIR /app

# COPY --from=base /app/bin/main .

# ENTRYPOINT ["./main"]