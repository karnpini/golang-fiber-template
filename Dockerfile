# Dockerfile (template)
FROM golang:1.24.5-alpine AS builder
RUN apk add --no-cache build-base git protobuf-dev protoc
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY . .
RUN chmod +x scripts/generate_proto.sh && sh scripts/generate_proto.sh
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o /app/main .

FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata curl
RUN curl -L https://github.com/fullstorydev/grpcurl/releases/download/v1.8.9/grpcurl_1.8.9_linux_x86_64.tar.gz | tar -xz -C /usr/local/bin grpcurl

WORKDIR /

COPY --from=builder /app/main .

CMD ["./main"]
