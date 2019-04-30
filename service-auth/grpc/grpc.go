package grpc

import (
	"google.golang.org/grpc"
	"os"
)

func ConnectUserService() (*grpc.ClientConn, error) {
	host := os.Getenv("USER_SERVICE_HOST")
	if len(host) == 0 {
		host = "localhost:11443"
	}

	return grpc.Dial(host)
}