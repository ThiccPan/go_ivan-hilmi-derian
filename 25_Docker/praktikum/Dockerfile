FROM golang:1.20 AS build

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build main.go
EXPOSE 8080
RUN ls

# stage 2
FROM alpine:3 as prod
WORKDIR /root/
COPY --from=build /usr/src/app/main .
COPY --from=build /usr/src/app/.env .
EXPOSE 8080
CMD ["./main"]