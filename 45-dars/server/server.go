package main

import (
	"context"
	"errors"
	"log"
	"mymode/proto/library"
	pb "mymode/proto/library"
	"mymode/storage"
	"mymode/storage/postgres"
	"net"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLibraryServiceServer
	b postgres.BookRepo
}

func main() {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := server{b: *postgres.NewBookRepo(db)}
	grpc := grpc.NewServer()
	pb.RegisterLibraryServiceServer(grpc, &s)

	if err = grpc.Serve(listener); err != nil {
		log.Fatal(err)
	}
	log.Printf("Server listening at %v...", listener.Addr())
}

func (S *server) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	err = S.b.CreateBook(library.Book{
		BookId:        id.String(),
		Title:         req.BookName,
		Author:        req.BookAuthor,
		YearPublished: req.YearPublished,
	})
	if err != nil {
		return nil, err
	}
	time.Sleep(2 * time.Second)
	return &pb.AddBookResponse{BookId: id.String()}, nil
}

func (s *server) SearchBook(ctx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	books, err := s.b.SearchBook(req.Query)
	if err != nil {
		return nil, err
	}
	time.Sleep(2 * time.Second)
	return &pb.SearchBookResponse{Books: books}, nil
}

func (s *server) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	status, err := s.b.GetByIdBook(req.BookId)
	if err != nil {
		return &pb.BorrowBookResponse{Status: false}, nil
	}
	if status == true {
		return &pb.BorrowBookResponse{Status: true}, nil
	}
	time.Sleep(2 * time.Second)
	return &pb.BorrowBookResponse{Status: false}, errors.New("Kutubxonamizda bu kitob mavjud emas.")
}
