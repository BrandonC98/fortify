FROM golang:latest AS BUILD_IMAGE

WORKDIR /src/app
COPY . .

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /usr/local/bin/interface ./...

FROM alpine:latest

ARG GIN_MODE=debug
ARG PASSMAN_PORT=8081
ARG PASSMAN_PASS_GEN_URL=http://localhost:8080
ARG PASSMAN_DATA_ACCESS_URL=http://localhost:8082

ENV GIN_MODE=$GIN_MODE
ENV PASSMAN_PORT=$PASSMAN_PORT
ENV PASSMAN_PASS_GEN_URL=$PASSMAN_PASS_GEN_URL
ENV PASSMAN_DATA_ACCESS_URL=$PASSMAN_DATA_ACCESS_URL

WORKDIR /passMan-interface

COPY --from=BUILD_IMAGE /usr/local/bin/interface /passMan-interface/interface
COPY static /passMan-interface/static
COPY templates /passMan-interface/templates

RUN chmod +x interface

EXPOSE $PASSMAN_PORT

CMD [ "./interface" ]
