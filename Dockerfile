FROM golang:1.20 as Builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 

WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod tidy

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o api .

FROM alpine:latest

WORKDIR /usr/local/bin

COPY --from=builder /build/api ./api

EXPOSE 7001

CMD ./api