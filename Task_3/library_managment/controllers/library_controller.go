package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strings"
)

func RegisterBook(l *services.Library) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the book title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter the book author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{
		Title:  title,
		Author: author,
		Status: "Available",
	}

	l.AddBook(book)
}

func RegisterMember(l *services.Library) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	member := models.Member{
		Name:          name,
		BorrowedBooks: []models.Book{},
	}
	l.AddMember(member)
}

func RemoveBook(l *services.Library) {
	var id int
	fmt.Print("Enter the book ID you want to remove: ")
	fmt.Scanln(&id)
	l.RemoveBook(id)
	fmt.Println("Book removed.")
}

func BorrowBook(l *services.Library) {
	var bookID, memberID int
	fmt.Print("Enter book ID to borrow: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter your member ID: ")
	fmt.Scanln(&memberID)

	err := l.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func ReturnBook(l *services.Library) {
	var bookID, memberID int
	fmt.Print("Enter book ID to return: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter your member ID: ")
	fmt.Scanln(&memberID)

	err := l.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func ListAvailableBooks(l *services.Library) {
	books := l.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("There are no available books.")
		return
	}

	fmt.Println("\nAvailable Books:")
	fmt.Printf("ID\tTitle\t\tAuthor\n")
	fmt.Println(strings.Repeat("-", 40))
	for _, book := range books {
		fmt.Printf("%d\t%s\t\t%s\n", book.ID, book.Title, book.Author)
	}
}

func ListBorrowedBooks(l *services.Library) {
	var memberID int
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)

	member, ok := l.Members[memberID]
	if !ok {
		fmt.Println("Member not found.")
		return
	}

	borrowed := l.ListBorrowedBooks(memberID)
	if len(borrowed) == 0 {
		fmt.Printf("%s has not borrowed any books.\n", member.Name)
		return
	}

	fmt.Printf("\nBooks borrowed by %s:\n", member.Name)
	fmt.Printf("ID\tTitle\t\tAuthor\n")
	fmt.Println(strings.Repeat("-", 40))
	for _, book := range borrowed {
		fmt.Printf("%d\t%s\t\t%s\n", book.ID, book.Title, book.Author)
	}
}
