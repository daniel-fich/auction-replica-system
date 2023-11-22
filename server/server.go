package main

import (
	. "auction-replica-system/shared"
	"fmt"
	"os"
	"time"

	"math/rand"
)

var (
	// HOSTNAME resolves to ip address
	dockerId   = os.Getenv("HOSTNAME")
	currentBid = MakeThreadVariable[int64](0)
)

func main() {
	server := Server{
		id: dockerId,
	}

	go startServer(server)
	time.Sleep(time.Second)
	WriteToSharedFile(dockerId, NodesFilename)

	for {
		time.Sleep(time.Second)
		randInt := rand.Intn(20)
		fmt.Println(randInt)
		if randInt == 0 {
			contents := GetFileContents("tmp", NodesFilename)
			removeDockerId := contents[0]

			RemoveDockerIdFromFile(removeDockerId, NodesFilename)
		}
		// go CheckPeers()
	}
}
