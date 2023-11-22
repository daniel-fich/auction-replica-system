package main

import (
	. "auction-replica-system/shared"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	// HOSTNAME resolves to ip address
	dockerId = os.Getenv("HOSTNAME")
)

func GetLeader() string {
	fileContents := GetFileContents(dockerId, NodesFilename)

	if len(fileContents) > 0 {
		return fileContents[0]
	}
	return ""
}

func randomBid(leader string) {
	result, err := getResult(leader)

	if err != nil {
		fmt.Println("Could not get result", err)
		return
	}

	// Our random bid will only bid one fourth of the time
	bid := rand.Int63n(4)
	if bid == 1 {
		newBid := result + rand.Int63n(15)
		fmt.Printf("Attempting to bid %d with current highest bid being %d\n", newBid, result)
		_, err := bidAmount(leader, newBid)

		if err != nil {
			fmt.Println("Failed to place bid: ", bid)
		}
	}
}

func main() {
	current := GetLeader()

	for {
		randomBid(current)
		time.Sleep(3 * time.Second)
	}
}
