package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Dinesh789kumar12/grpc-go/greet"
	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("0.0.0.0:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := greet.NewGreetServiceClient(con)
	req := greet.GreetRequest{ReqGreet: "Dinesh"}
	res, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
