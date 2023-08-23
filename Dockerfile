# Use the official Golang image as a builder.
FROM golang:1.20.6 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# Allows for caching of dependencies unless go.{mod,sum} change.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build each microservice and the API gateway.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o comment ./service/comment
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o favorite ./service/favorite
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o feed ./service/feed
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o message ./service/message
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o pong ./service/pong
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o publish ./service/publish
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o relation ./service/relation
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o user ./service/user
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o api-gateway ./api

# Use the Alpine image for the runtime container.
FROM alpine:3.14
RUN apk --no-cache add ca-certificates

# Copy each compiled binary from the builder stage.
COPY --from=builder /app/comment /app/comment
COPY --from=builder /app/favorite /app/favorite
COPY --from=builder /app/feed /app/feed
COPY --from=builder /app/message /app/message
COPY --from=builder /app/pong /app/pong
COPY --from=builder /app/publish /app/publish
COPY --from=builder /app/relation /app/relation
COPY --from=builder /app/user /app/user
COPY --from=builder /app/api-gateway /app/api-gateway

# Add a script to start services and the API gateway in the specified order.
COPY start-services.sh /app/start-services.sh
RUN chmod +x /app/start-services.sh

CMD ["/app/start-services.sh"]
