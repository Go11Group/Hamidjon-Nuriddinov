package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "mymode/proto/library"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewLibraryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	Menu(c, ctx)
}

func Menu(c pb.LibraryServiceClient, ctx context.Context) {
	fmt.Println(`
		1 - Add Book
		2 - Search Book
		3 - Borrow Book
		4 - Exit`)
	var buyruq int
	fmt.Scan(&buyruq)
	switch buyruq {
	case 1:
		AddBook(c, ctx)
	case 2:
		SearchBook(c, ctx)
	case 3:
		BorrowBook(c, ctx)
	case 4:
		panic(0)
	}
}

func AddBook(c pb.LibraryServiceClient, ctx context.Context) {
	book := pb.AddBookRequest{
		BookName:      "1984",
		BookAuthor:    "J.Oruel",
		YearPublished: 1947,
	}
	r, err := c.AddBook(ctx, &book)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nKitob kutubxonaga qo'shildi. Id: %v", r.BookId)
	Menu(c, ctx)
}

func SearchBook(c pb.LibraryServiceClient, ctx context.Context) {
	query := pb.SearchBookRequest{
		Query: "1984",
	}
	r, err := c.SearchBook(ctx, &query)
	if err != nil {
		log.Fatal(err)
	}
	for _, book := range r.Books {
		log.Println(book)
	}
	Menu(c, ctx)
}

func BorrowBook(c pb.LibraryServiceClient, ctx context.Context) {
	req := pb.BorrowBookRequest{
		UserId: "123",
		BookId: "73b7839f-338d-11ef-9d6c-5c60baca83e2",
	}
	r, err := c.BorrowBook(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: %v", r.Status)
	Menu(c, ctx)
}
