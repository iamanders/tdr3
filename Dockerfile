FROM golang:1.20.5 as builder
# FROM --platform=linux/amd64 golang:1.20.5 as builder

WORKDIR /app/

# COPY go.mod, go.sum and install the dependencies
COPY go.* ./
RUN go mod download

# Copy all files
COPY . .

RUN CGO_ENABLED=1 go build -o /app/tdr3 -a -ldflags '-linkmode external -extldflags "-static"' ./cmd/web/
# RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /app/tdr3 -a -ldflags '-linkmode external -extldflags "-static"' ./cmd/web/

FROM alpine:latest
# FROM scratch

COPY --from=builder /app/data /app/data
COPY --from=builder /app/frontend /app/frontend
COPY --from=builder /app/tdr3 /app/tdr3

EXPOSE 8888

WORKDIR /app

ENTRYPOINT ["/app/tdr3"]
