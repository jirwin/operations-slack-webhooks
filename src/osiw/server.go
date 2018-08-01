package osiw

import (
	"net"

	"bytes"
	"net/http"

	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	listenAddr string
	webhookUrl string
	server     *grpc.Server
}

func (s *Server) Post(ctx context.Context, req *PostRequest) (*PostReply, error) {
	payload, err := req.GetPayload()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	r, err := http.NewRequest("POST", s.webhookUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	defer resp.Body.Close()

	return &PostReply{}, nil
}

func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	RegisterOswiServer(grpcServer, s)
	s.server = grpcServer

	fmt.Printf("Starting osiw server on %s\n", lis.Addr())
	go grpcServer.Serve(lis)

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.server.GracefulStop()
	return nil
}

func NewServer(listenAddr, webhookUrl string) *Server {
	return &Server{
		listenAddr: listenAddr,
		webhookUrl: webhookUrl,
	}
}
