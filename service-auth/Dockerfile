FROM golang:1.12-alpine AS builder

EXPOSE 8080

#disable crosscompiling
ENV CGO_ENABLED=0

#enable go mod
ENV GO111MODULE=on

#compile linux only
ENV GOOS=linux

# Git required for fetching the dependencies.
RUN apk update && apk add --no-cache git ca-certificates tzdata

# Installing
WORKDIR $GOPATH/src/github.com/deissh/api.micro

COPY go.mod go.mod
COPY go.sum go.sum

# Using go mod
RUN go mod download

# Copy file(s)
COPY . ./

WORKDIR $GOPATH/src/github.com/deissh/api.micro/service-auth

# Build the binary - remove debug info and compile only for linux target
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o /go/bin/service .

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy static executable
COPY --from=builder /go/bin/service /service
ENTRYPOINT ["/service"]