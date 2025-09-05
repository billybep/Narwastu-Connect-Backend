FROM golang:1.25 AS builder

WORKDIR /src

# copy dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build binary
RUN go build -o /bin/server ./cmd/server

# Stage 2: final image (lebih ringan)
FROM alpine:3.20

WORKDIR /app

# copy hasil build dari stage builder
COPY --from=builder /bin/server /app/server

EXPOSE 8080
CMD ["/app/server"]
