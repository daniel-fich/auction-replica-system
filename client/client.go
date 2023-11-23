package main

import (
	proto "auction-replica-system/grpc"
	. "auction-replica-system/shared"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	// HOSTNAME resolves to ip address
	dockerId = os.Getenv("HOSTNAME")
)

func randomBid(leader string) {
	result, err := GetResult(leader)

	if err != nil {
		log.Println("Could not get result", err)
		return
	}

	if result.OutcomeType == proto.OutcomeType_RESULT {
		log.Printf("Result has been recieved outcome is %d\n", result.Outcome)
		return
	}

	// Our random bid will only bid one fourth of the time
	bid := rand.Int63n(4)
	if bid == 1 {
		newBid := result.Outcome + rand.Int63n(15)
		log.Printf("Attempting to bid %d with current highest bid being %d to leader %s\n", newBid, result, leader)
		_, err := BidAmount(leader, newBid)

		if err != nil {
			log.Printf("Failed to place bid: %d", newBid)
		}
	}
}

func main() {
	for {
		current := GetLeader(dockerId)
		randomBid(current)
		time.Sleep(3 * time.Second)
	}
}
