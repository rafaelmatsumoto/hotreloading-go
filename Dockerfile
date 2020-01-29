FROM golang:1.13-alpine as base
RUN apk add git
EXPOSE 8000

FROM base as dev
RUN go get github.com/cespare/reflex
COPY reflex.conf /
ENTRYPOINT ["reflex", "-c", "/reflex.conf"]
