package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/arijitnayak92/go-lang-projects/BIGRPC/proto"
	"github.com/twinj/uuid"
)

type NotificationService interface {
	PushNotification() (string, error)
}

type NotificationCient struct {
	client pb.NotificationClient
}

func NewNotificationClient(client pb.NotificationClient) *NotificationCient {
	return &NotificationCient{client: client}
}

func (n *NotificationCient) PushNotification() (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := n.client.PushNotification(ctx)
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	waitResponse := make(chan error)

	// first goroutine sends random increasing numbers to stream
	// and closes int after 10 iterations
	go func() {
		for i := 1; i <= 10; i++ {
			// generate random nummber and send it to stream

			req := pb.NotificationRequest{
				UserID:   uuid.NewV4().String(),
				TargetID: uuid.NewV4().String(),
				Message:  "new notification",
			}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			time.Sleep(time.Millisecond * 200)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Println(err)
		}

		waitResponse <- err

	}()

	// second goroutine receives data from stream
	// and saves result in max variable
	//
	// if stream is finished it closes done channel
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("cannot receive stream response: %v", err)
				return
			}

			log.Println("new received", res.NotificationID)
		}
	}()

	// third goroutine closes done channel
	// if context is done
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(waitResponse)
	}()

	err = <-waitResponse
	log.Println("done !")
	return "", err
}
