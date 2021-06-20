package main

import (
	v1 "04-project-layout/api/greeter/v1"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "127.0.0.1:50001"
)


func addGreeter(conn *grpc.ClientConn){

	c := v1.NewGreeterClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(), time.Second)
	defer  cancel()
	req := &v1.CreateGreeterRequest{Job: "IT", Word: "我是一个好孩子"}
	r, err := c.CreateGreeter(ctx,req)
	if err != nil {
		log.Fatalf("not add greeter: %v", err)
	}
	log.Printf("greeter add: %s", r.String())
}

func updateGreeter(conn *grpc.ClientConn){

	c := v1.NewGreeterClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(), time.Second)
	defer  cancel()
	req := &v1.UpdateGreeterRequest{Id: 1,Job: "老师", Word: "我是一个好老师"}
	r, err := c.UpdateGreeter(ctx,req)
	if err != nil {
		log.Fatalf("not update greeter: %v", err)
	}
	log.Printf("greeter update: %s", r.String())
}


func getGreeter(conn *grpc.ClientConn){

	c := v1.NewGreeterClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(), time.Second)
	defer  cancel()
	req := &v1.GetGreeterRequest{Id:1}
	r, err := c.GetGreeter(ctx,req)
	if err != nil {
		log.Fatalf("not get greeter: %v", err)
	}
	log.Printf("greeter get: %s", r.String())
}


func listGreeter(conn *grpc.ClientConn){

	c := v1.NewGreeterClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(), time.Second)
	defer  cancel()
	req := &v1.ListGreeterRequest{}
	r, err := c.ListGreeter(ctx,req)
	if err != nil {
		log.Fatalf("not list greeter: %v", err)
	}
	log.Printf("greeter list: %s", r.String())
}


func deleteGreeter(conn *grpc.ClientConn){

	c := v1.NewGreeterClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(), time.Second)
	defer  cancel()
	req := &v1.DeleteGreeterRequest{Id: 2}
	r, err := c.DeleteGreeter(ctx,req)
	if err != nil {
		log.Fatalf("not delete greeter: %v", err)
	}
	log.Printf("greeter delete: %s", r.String())
}

func main(){

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connet: %v", err)
	}

	defer  conn.Close()

	//addGreeter(conn)
	//updateGreeter(conn)
	//getGreeter(conn)
	//listGreeter(conn)
	//deleteGreeter(conn)
	listGreeter(conn)
}