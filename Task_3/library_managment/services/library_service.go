package services

import (
	"errors"
	"fmt"
	"library_management/models"
)
type libraryManager interface{
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book

}
type Library struct{
	Books map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books: make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	book.ID = len(l.Books) + 1
	l.Books[book.ID] = book 
	fmt.Printf("Book registered successfully! Book ID is %d\n", book.ID)
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("book with id: %d not found", bookID)
	}
	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("member with id: %d not found", memberID)
	}
	if book.Status == "Borrowed"{
		return errors.New("book is already borrowed")
	}
	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	return nil

}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, ok :=l.Books[bookID]
	if !ok {
		return fmt.Errorf("book with id: %d not found", bookID)
	}
	member, ok := l.Members[memberID]

	if !ok {
		return fmt.Errorf("member with id: %d not found ", memberID)
	}

	found := false
	newBorrowed := []models.Book{}

	for _, b := range member.BorrowedBooks {
		if b.ID == bookID{
			found = true
		} else{
			newBorrowed = append(newBorrowed, b)
		}
	}
	if !found{
		return errors.New("book not borrowed by this member")
	}

	book.Status = "Available"
	l.Books[bookID] = book
	member.BorrowedBooks = newBorrowed
	l.Members[memberID] = member
	return nil

}

func (l *Library) ListAvailableBooks() []models.Book {
	available := []models.Book{}
	for _, book := range l.Books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	if member, ok := l.Members[memberID];ok {
		return member.BorrowedBooks
	}
	return nil
}

func (l *Library) AddMember(member models.Member) {
	member.ID = len(l.Members) + 1
	l.Members[member.ID] = member
	fmt.Printf("Member registered successfully! Your ID is %d\n", member.ID)
}

