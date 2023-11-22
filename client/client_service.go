package main

import (
	proto "auction-replica-system/grpc"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const SERVER_PORT = 6969

func bidAmount(serverId string, amount int64) (proto.ErrorCode, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", serverId, strconv.Itoa(SERVER_PORT)),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("Could not connect to server: ", serverId)
		return 0, errors.New("Could not connect to server: " + serverId)
	}

	auctionClient := proto.NewAuctionServiceClient(conn)

	ack, err := auctionClient.Bid(context.Background(), &proto.Amount{
		Amount: amount,
	})

	if err != nil {
		log.Println("Could not place bid on server: ", serverId)
		return 0, errors.New("Could not place bid on server: " + serverId)
	}

	return ack.ErrorCode, nil
}

func getResult(serverId string) (int64, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", serverId, strconv.Itoa(SERVER_PORT)),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("Could not connect to server: ", serverId)
		return 0, errors.New("Could not connect to server: " + serverId)
	}

	auctionClient := proto.NewAuctionServiceClient(conn)

	outcome, err := auctionClient.Result(context.Background(), &proto.EmptyArgument{})

	if err != nil {
		log.Println("Could not get result from server: ", serverId, " ", err)
		return 0, errors.New("Could not get result from server: " + serverId)
	}

	return outcome.Outcome, nil
}
