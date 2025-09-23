package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Pages  int    `json:"pages"`
}

func readBookInfo(filename string) (Book, error) {
	var book Book

	data, err := os.ReadFile(filename)
	if err != nil {
		return book, fmt.Errorf("Ошибка чтения файла: %w", err)
	}

	err = json.Unmarshal(data, &book)
	if err != nil {
		return book, fmt.Errorf("Ошибка парсинга JSON: %w", err)
	}

	return book, nil

}

func printBookInfo(book Book) {
	fmt.Printf("Книга: %s, Автор: %s, Год: %d, Страниц: %d\n", book.Title, book.Author, book.Year, book.Pages)
}

func main() {

	book, err := readBookInfo("book.json")
	if err != nil {
		log.Fatalf("Ошибка: %v", err)

	}
	printBookInfo(book)

}
