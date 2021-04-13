package grpchandler

import "github.com/arijitnayak92/go-lang-projects/BIGRPC/domain"

type GrpcHandler struct {
	domain *domain.AMSNotification
}

func NewGrpcHandler(domain *domain.AMSNotification) *GrpcHandler {
	return &GrpcHandler{domain: domain}
}
