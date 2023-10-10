package stream

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	pb "monit/pb/client/stream"
	"monit/pkg/common/config"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type videoServiceClient struct {
	Client pb.VideoServiceClient
}

func InitVideoClient(c config.Config) (VideoServiceClient, error) {
	fmt.Println("step 2")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	caCert, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate: %w", err)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, errors.New("failed to append CA certificate to certificate pool")
	}

	//read client cert
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, fmt.Errorf("failed to load client certificate and key: %w", err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	tlsCredential := credentials.NewTLS(config)
	fmt.Println(c.StreamUrl)
	clientCon, err := grpc.DialContext(ctx, c.StreamUrl, grpc.WithTransportCredentials(tlsCredential))
	if err != nil {
		return nil, fmt.Errorf("failed to load client certificate and key: %w", err)
	}
	fmt.Println("step 3")

	return NewConferenceClient(pb.NewVideoServiceClient(clientCon)), nil
}

func NewConferenceClient(client pb.VideoServiceClient) VideoServiceClient {
	fmt.Println("step 4")

	return &videoServiceClient{
		Client: client,
	}
}

func (c *videoServiceClient) HealthCheck(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	resp, err := c.Client.HealthCheck(ctx, &pb.Request{
		Data: "Hi authentication server",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *videoServiceClient) VideoDetails(ctx context.Context, req *pb.VideoDetailsRequest) (*pb.VideoDetailsResponse, error) {
	fmt.Println("step 1")
	resp, err := c.Client.VideoDetails(ctx, &pb.VideoDetailsRequest{VideoID: req.VideoID})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}
