FROM golang:1.12-alpine AS builder

EXPOSE 8000

# Git required for fetching the dependencies.
RUN apk update && apk add --no-cache git ca-certificates tzdata

# Create appuser. Not using since rev-proxy currently on :80 (< :1024 req. root)
# RUN adduser -D -g '' appuser

# Copy file(s)
WORKDIR /go/src/service
COPY main.go .

# Using vgo
RUN go mod download

#disable crosscompiling
ENV CGO_ENABLED=0

#compile linux only
ENV GOOS=linux

# Build the binary - remove debug info and compile only for linux target
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o /go/bin/service .
# RUN go build -o /go/bin/app

############################
# STEP 2 build a small image
############################
FROM scratch

# COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy static executable
COPY --from=builder /go/bin/service /service

# Use an unprivileged user.
# USER appuser

# Run the hello binary.
ENTRYPOINT ["/service"]