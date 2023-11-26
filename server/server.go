package main

import (
	. "auction-replica-system/shared"
	"errors"
	"log"
	"os"
	"time"
)

var (
	// HOSTNAME resolves to ip address
	dockerId      = os.Getenv("HOSTNAME")
	currentBid    = MakeThreadVariable[int64](0)
	currentLeader = ""
	isAuctionOver = MakeThreadVariable[bool](false)
)

func randomCrash() {
	if leader, err := GetCurrentLeader(); err == nil {
		if leader == dockerId {
			time.Sleep(5 * time.Second)
			log.Fatal("Oh no! Something awful happened to my server, I spilled coffee all over it!")
		}
	}
	// for {
	// 	time.Sleep(5 * time.Second)
	// 	shouldExit := rand.Int63n(10)

	// 	if shouldExit == 1 {
	// 		log.Fatal("Oh no! Something awful happened to my server, I spilled coffee all over it!")
	// 	}
	// }
}

func GetCurrentLeader() (string, error) {
	nodeIds := GetFileContents(dockerId, NodesFilename)
	if len(nodeIds) == 0 {
		return "", errors.New("No ids found")
	}
	return nodeIds[0], nil
}

func main() {
	server := Server{
		id: dockerId,
	}

	go startServer(server)
	WriteToSharedFile(dockerId, NodesFilename)
	time.Sleep(time.Second)

	go CheckPeers()
	go synchronize()
	go startTimeServer()
	time.Sleep(time.Second)
	SetExpiration()
	go PollTime()
	go randomCrash()
	for {
	}
}
