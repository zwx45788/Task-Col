package user

import (
	"log"
	loginServiceV1 "project-user/pkg/service/login.service.v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var LoginServiceClient loginServiceV1.LoginServiceClient

func InitUserRpc() {
	conn, err := grpc.NewClient("etcd:///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	LoginServiceClient = loginServiceV1.NewLoginServiceClient(conn)
	loginServiceV1.NewLoginServiceClient(conn)
	//c := pb.NewGreeterClient(conn)
	//
	//// Contact the server and print out its response.
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetMessage())
}
