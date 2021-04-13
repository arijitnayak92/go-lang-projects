package client

import (
	"context"
	"io"
	"log"
	"math/rand"
	"time"

	pb "github.com/arijitnayak92/go-lang-projects/BIGRPC/proto"
	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().Unix())

	// dail server
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := pb.NewNotificationClient(conn)
	stream, err := client.PushNotification(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	var max int32
	ctx := stream.Context()
	done := make(chan bool)

	// first goroutine sends random increasing numbers to stream
	// and closes int after 10 iterations
	go func() {
		for i := 1; i <= 10; i++ {
			// generate random nummber and send it to stream
			rnd := int32(rand.Intn(i))
			rnd2 := int32(rand.Intn(i))

			req := pb.NotificationRequest{
				UserID:   string(rnd),
				TargetID: string(rnd2),
				Message:  "new notification",
			}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			time.Sleep(time.Millisecond * 200)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	// second goroutine receives data from stream
	// and saves result in max variable
	//
	// if stream is finished it closes done channel
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			log.Println("new received", resp.NotificationID)
		}
	}()

	// third goroutine closes done channel
	// if context is done
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
	log.Printf("finished with max=%d", max)
}
