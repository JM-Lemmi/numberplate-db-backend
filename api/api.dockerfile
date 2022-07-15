FROM golang:alpine AS build

RUN apk --update add git

WORKDIR /app
COPY . /app

RUN go build .

FROM alpine AS run

COPY --from=build /app/numberplatedb-api /app/numberplatedb-api

WORKDIR /app
EXPOSE 80

CMD ./numberplatedb-api
