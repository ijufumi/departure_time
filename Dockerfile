FROM golang:1.19-bullseye as build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main /app/cmd/app/

WORKDIR /app/web

RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
RUN apt-get update \
&& apt-get install -y yarn nodejs npm \
&& npm install npx -g \
&& yarn build \
&& rm -rf /var/cache/apt

FROM debian:bullseye-slim as deploy

WORKDIR /app

RUN apt-get update \
&& apt-get install -y ca-certificates \
&& rm -rf /var/cache/apt

COPY --from=build /app/main /app/main

COPY --from=build /app/web/public /app/web/public

EXPOSE 8080

ENTRYPOINT ["/app/main"]
