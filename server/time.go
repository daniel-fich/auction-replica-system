package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	. "auction-replica-system/shared"
	proto "auction-replica-system/time"

	"github.com/beevik/ntp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	expiration      time.Time
	isExpirationSet bool = false
)

const TIME_SERVER_PORT = 420

type TimeService struct {
	proto.UnimplementedTimeServiceServer
}

// Sets the expiration time to 30 seconds in the future based on time queried from ntp.org
func SetExpiration() {
	if leader, err := GetCurrentLeader(); err == nil {
		if leader == dockerId {
			response, err := ntp.Query("0.dk.pool.ntp.org")
			if err != nil {
				log.Fatal("Failed to get end time for auction. Server shutting down...")
			}
			expiration = response.Time.Add(time.Second * 10)
			isExpirationSet = true
			for _, id := range GetFileContents(dockerId, NodesFilename) {
				if id != dockerId {
					SetTime(id)
				}
			}
		}
	}
}

func SetTime(serverId string) error {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", serverId, strconv.Itoa(TIME_SERVER_PORT)),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("Could not connect to server: ", serverId)
		return errors.New("Could not connect to server: " + serverId)
	}

	auctionClient := proto.NewTimeServiceClient(conn)

	_, err = auctionClient.SetTime(context.Background(), &proto.Time{
		Time: uint64(expiration.Unix()),
	})

	if err != nil {
		log.Println("Could not place bid on server: ", serverId)
		return errors.New("Could not place bid on server: " + serverId)
	}

	return nil
}

func startTimeServer(server TimeService) {
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(TIME_SERVER_PORT))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}

	log.Printf("Started server at port: %d\n", TIME_SERVER_PORT)

	// Register the grpc server and serve its listener
	proto.RegisterTimeServiceServer(grpcServer, &server)
	serveError := grpcServer.Serve(listener)

	if serveError != nil {
		log.Fatal("Could not serve listener")
	}
}

func (t *TimeService) SetTime(ctx context.Context, in *proto.Time) (*proto.EmptyMessage, error) {
	expiration = time.Unix(int64(in.Time), 0)
	isExpirationSet = true
	log.Println("Time has been set to", expiration.Format("hh:mm:ss"))
	return &proto.EmptyMessage{}, nil
}

func PollTime() {
	for {
		time.Sleep(time.Millisecond * 100)
		if isExpirationSet {
			response, err := ntp.Query("0.dk.pool.ntp.org")
			if err != nil {
				log.Fatal("Failed to get time from ntp server. Server shutting down...")
			}
			if response.Time.After(expiration) {
				isAuctionOver.Update(true)
			}
		}
	}
}
