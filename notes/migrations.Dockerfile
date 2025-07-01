FROM golang:1.24-alpine
RUN apt update
RUN GOBIN=/ go install github.com/rubenv/sql-migrate/...@latest
COPY ./internal/schema/ .
RUN ["sql-migrate up", "- config=.dbconfig.yaml", "-env=docker"]
