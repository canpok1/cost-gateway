# === ビルド用 ===
FROM golang:1.23-bookworm AS builder
WORKDIR /app
COPY . .

RUN make build

# === 実行用 ===
FROM debian:bookworm-slim
WORKDIR /app

COPY --from=builder /app/build/release/server ./

ENV TZ=Asia/Tokyo
ENV MYSQL_DATABASE=database
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=password
ENV MYSQL_HOST=localhost
ENV MYSQL_PORT=3306

EXPOSE 8080

ENTRYPOINT [ "/app/server" ]
