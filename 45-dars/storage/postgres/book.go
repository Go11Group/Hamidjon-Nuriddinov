package postgres

import (
	"database/sql"
	"mymode/proto/library"
)

type BookRepo struct {
	Db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{Db: db}
}

func (B *BookRepo) CreateBook(book library.Book) error {
	_, err := B.Db.Exec(`INSERT INTO 
							Books(id, name, author, year) 
						values($1, $2, $3, $4)`,
		book.BookId, book.Title, book.Author, book.YearPublished)
	return err
}

func (B *BookRepo) SearchBook(query string) ([]*library.Book, error) {
	books := []*library.Book{}
	rows, err := B.Db.Query(`SELECT 
							*
						FROM 
							Books
						WHERE
							name = $1
						OR
							author = $1`, query)
	if err != nil {
		return books, nil
	}
	for rows.Next() {
		book := library.Book{}
		err = rows.Scan(&book.BookId, &book.Title, &book.Author, &book.YearPublished)
		if err != nil {
			return books, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func (B *BookRepo) GetByIdBook(bookId string) (bool, error) {
	book := library.Book{}
	err := B.Db.QueryRow(`SELECT * FROM Books WHERE id = $1`, bookId).Scan(&book.BookId, &book.Title, &book.Author, &book.YearPublished)
	if err != nil || book.BookId == "" {
		return false, err
	}
	return true, nil
}
