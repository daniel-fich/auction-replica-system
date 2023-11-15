package main

import (
	. "auction-replica-system/shared"
	"os"
)

var (
	// Resolves to ip address
	dockerId = os.Getenv("HOSTNAME")
)

func main() {
	server := Server{
		id: dockerId,
	}

	go startServer(server)
	WriteToSharedFile(dockerId, "hosts")
}
