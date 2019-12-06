# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git make
ADD . /src
RUN cd /src && make build

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/bin/microbe /app/microbe
ENV GIN_MODE release
ENTRYPOINT ./microbe