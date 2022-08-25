FROM golang:1.19-bullseys as build

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o cmd/app/main.go

FROM debian:bullseye-slim as deploy

WORKDIR /app

COPY --from=build /app/main /app/main

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/main"]