FROM golang:1.20-bullseye as builder

COPY . /app
WORKDIR /app
EXPOSE 8000
RUN go build -o main pkg/main/main.go

FROM golang:1.20-bullseye as runner
COPY --from=builder /app/main /app/main

ENV COLLECTION_NAME=nope
ENV LEDGER=nope
ENV API_KEY=nope
ENV VAULT_URL=https://vault.immudb.io/ics/api/v1
ENV CERT_FILE=""
ENV KEY_FILE=""
ENV BIND_ADDR="0.0.0.0:8080"


ENTRYPOINT ["/app/main"]