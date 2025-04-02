FROM golang:1.24 AS base

WORKDIR /app

COPY .env .
COPY scripts/build.sh .
COPY cmd/ cmd/
COPY infra/ infra/
COPY src/ src/
COPY main.go .

RUN apt-get -y update
RUN go mod init Go-Hexagonal
RUN go mod tidy
RUN chmod +x ./build.sh
RUN ./build.sh web main

ENTRYPOINT ["bin/main"]


# TODO: fix "exec ./main: no such file or directory"
# FROM scratch

# WORKDIR /app

# COPY --from=base /app/bin/main .

# ENTRYPOINT ["./main"]