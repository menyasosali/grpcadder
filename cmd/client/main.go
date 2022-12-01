package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/menyasosali/grpcadder/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"
	"strconv"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to the server: %s", err)
	}

	c := api.NewAdderClient(conn)
	res, err := c.Add(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.GetResult())
}
