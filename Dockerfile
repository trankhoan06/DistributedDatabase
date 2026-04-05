#stage build
#Build image base on base image
FROM golang:1.23-alpine as builder
COPY ./ /app/
WORKDIR /app/
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o esim .

#stage runner
FROM alpine
WORKDIR /app/
COPY --from=builder /app/esim .
COPY .env .env
#COPY secret secret
COPY wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh
#COPY db/migration db/migration
#CMD ["make migrate_up"]
CMD ["/app/esim"]