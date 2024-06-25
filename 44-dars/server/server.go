package main

import (
	"context"
	"log"
	pb "mymode/proto"
	"net"
	"strings"
	"unicode"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTranslateServiceServer
}

func (s *server) TranslateText(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateRespons, error) {
	var translatedText string
	for _, i := range req.Words {
		var word string
		switch strings.ToLower(i) {
		case "hello":
			word = "salom"
		case "world":
			word = "dunyo"
		case "apple":
			word = "olma"
		default:
			word = "golang"
		}
		translatedText += " " + word
	}
	return &pb.TranslateRespons{TranslatedText: strings.TrimLeftFunc(translatedText, unicode.IsSpace)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslateServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
