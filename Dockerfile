FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /go/bin/app -ldflags '-extldflags "-static"' -tags timetzdata

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app"]
