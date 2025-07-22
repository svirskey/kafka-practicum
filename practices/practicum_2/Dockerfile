# BUILD STAGE
FROM golang:latest as builder

WORKDIR /application

COPY ./application/ .

RUN go mod download && go mod verify

RUN go build -o application

# RUN STAGE
FROM ubuntu:24.04

WORKDIR /application

COPY --from=builder /application/application .

CMD ["/application/application", "-mode=release"]