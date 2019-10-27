FROM golang:latest
RUN mkdir -p /work
WORKDIR /work
ADD ./go ./go
