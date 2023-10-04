run:
	go run cmd/main.go
proto:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH=$$PATH:$$(go env GOPATH)/bin
	protoc -I ./pb ./pb/server/monit.proto --go_out=. --go-grpc_out=.
docker:
	sudo docker build -t edwinsiby/monitization-server .