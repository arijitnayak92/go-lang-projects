package server

import (
	"io"
	"log"
	"sync"

	pb "github.com/arijitnayak92/go-lang-projects/BIGRPC/proto"
)

type NotificationServer struct{}

type NotificationStore struct {
	notifications []Notification
	sync.Mutex
}

func NewNotificationStore() *NotificationStore {
	return &NotificationStore{}
}

func (n *NotificationServer) PushNotification(srv pb.Notification_PushNotificationServer) error {
	ctx := srv.Context()

	for {
		// exit on context completion
		// or else continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// close the stream
			log.Println("end of the stream")
			return nil
		}
		if err != nil {
			log.Println("error while receiving stream data : ", err)
			continue
		}

		notificationReq := Notification{
			UserID:  req.GetUserID(),
			RefID:   req.GetTargetID(),
			Message: req.GetMessage(),
		}
		notificationStore := NewNotificationStore()

		var wg sync.WaitGroup
		wg.Add(1)

		go func(notificationReq Notification) {
			defer wg.Done()

			notificationStore.Lock()
			notificationStore.notifications = append(notificationStore.notifications, notificationReq)
			notificationStore.Unlock()

		}(notificationReq)

		wg.Wait()

		resp := pb.NotificationResponse{
			NotificationID: "",
		}

		if err := srv.Send(&resp); err != nil {
			log.Println("error while sending response : ", err)
		}
		log.Println("store data : ", notificationStore.notifications)
		log.Println("sent response !")
	}
}
