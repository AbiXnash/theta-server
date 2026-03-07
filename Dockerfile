# syntax=docker/dockerfile:1

FROM golang:1.25.7-alpine AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags='-s -w' -o /out/theta-server .

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app

COPY --from=builder /out/theta-server /app/theta-server
COPY data /app/data

ENV PORT=8080
ENV THETA_DATA_DIR=/app/data

EXPOSE 8080

ENTRYPOINT ["/app/theta-server"]
