# syntax=docker/dockerfile:1
FROM golang:1.23 AS builder
WORKDIR /src
COPY . .
RUN go build -o /bin/depscheck ./cmd/

FROM gcr.io/distroless/static-debian12
COPY --from=builder /bin/depscheck /depscheck
CMD ["/depscheck"]

