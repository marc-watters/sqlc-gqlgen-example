FROM golang:1.24.1-alpine3.21 AS base
WORKDIR /api/
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM base AS build
COPY . .
RUN go build -o sqlc-gqlgen-example ./cmd/

FROM build AS run
ARG API_HOST
ARG API_PORT
ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PASS
ARG DB_NAME

ENV API_HOST=${API_HOST}
ENV API_PORT=${API_PORT}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_USER=${DB_USER}
ENV DB_PASS=${DB_PASS}
ENV DB_NAME=${DB_NAME}

ENTRYPOINT [ "./sqlc-gqlgen-example" ]
