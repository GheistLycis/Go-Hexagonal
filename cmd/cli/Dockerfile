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
RUN chmod +x ./build.sh
RUN ./build.sh cli main


# TODO: fix "exec ./main: no such file or directory"
# FROM scratch

# WORKDIR /app

# COPY --from=base /app/bin/main .