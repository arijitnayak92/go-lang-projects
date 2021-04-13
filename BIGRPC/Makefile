gen:
	rm -f proto/*.pb.go
	protoc -I proto proto/*.proto --go_out=plugins=grpc:proto/
all: 
	client server
server: protoc
	@echo "Building server"
	go build -o server \
		github.com/arijitnayak92/go-lang-projects/BIGRPC/server
client: protoc
	@echo "Building client"
	go build -o client \
		github.com/arijitnayak92/go-lang-projects/BIGRPC/client