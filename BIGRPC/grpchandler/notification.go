package grpchandler

import (
	"io"
	"log"

	pb "github.com/arijitnayak92/go-lang-projects/BIGRPC/proto"
)

func (n GrpcHandler) PushNotification(srv pb.Notification_PushNotificationServer) error {
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

		createdID, err := n.domain.AddNotification(req.GetUserID(), req.GetTargetID(), req.GetMessage())

		if err != nil {
			log.Println("failed to add notification : ", err)
			return err
		}
		resp := pb.NotificationResponse{
			NotificationID: createdID,
		}

		if err := srv.Send(&resp); err != nil {
			log.Println("error while sending response : ", err)
		}

		log.Println("sent response !")

	}

}
