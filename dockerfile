FROM golang:latest AS BUILD_IMAGE

WORKDIR /src/app
COPY . .

RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /usr/local/bin/pass-gen ./...

FROM alpine:latest

ARG GIN_MODE=release
ARG PASSMAN_PORT=8080
ARG PASSMAN_STRING_MIN=7
ARG PASSMAN_STRING_MAX=25

ENV GIN_MODE=$GIN_MODE
ENV PASSMAN_PORT=$PASSMAN_PORT
ENV PASSMAN_STRING_MIN=$PASSMAN_STRING_MIN
ENV PASSMAN_STRING_MAX=$PASSMAN_STRING_MAX

COPY --from=BUILD_IMAGE /usr/local/bin/pass-gen /usr/bin/pass-gen

EXPOSE $PASSMAN_PORT

CMD [ "pass-gen" ]
