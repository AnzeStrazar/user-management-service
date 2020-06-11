FROM golang:1.14 AS build
WORKDIR /src
# Copy go.mod and go.sum to download all dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy everything else to actually build the project
COPY . .
# Copy config.json to build folder
COPY config/config.json build/config/config.json
# Build project to build folder
RUN go build -o build/user-management-service

FROM debian:stable-slim
COPY --from=build /src/build/ /app/
WORKDIR /app
ENTRYPOINT ["/app/user-management-service"]