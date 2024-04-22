package main

import (
	"context"
	"encoding/json"
	"log"

	firestore "cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func SaveBookData(u *Book) error {
	_, _, err := Client.Collection("users").Doc(userData.Email).Collection("books").Add(context.Background(), map[string]interface{}{
		"title":      u.Title,
		"page_total": u.PageTotal,
		"written_by": u.WrittenBy,
	})

	return err
}

func ReadBookData() ([]Book, error) {
	iter := Client.Collection("users").Doc(userData.Email).Collection("books").Documents(context.Background())
	var books []Book
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return []Book{}, err
		}
		jsonData, _ := json.Marshal(doc.Data())

		var book Book
		json.Unmarshal(jsonData, &book)
		book.DocId = doc.Ref.ID
		books = append(books, book)
		log.Println(book.Title)

	}

	return books, nil
}

func UpdateBookData(docId string, u *Book) error {
	_, err := Client.Collection("users").Doc(userData.Email).Collection("books").Doc(docId).Update(context.Background(), []firestore.Update{
		{
			Path:  "title",
			Value: u.Title,
		},
		{
			Path:  "page_total",
			Value: u.PageTotal,
		},
		{
			Path:  "written_by",
			Value: u.WrittenBy,
		},
	})

	return err
}

func DeleteBookData(docId string) error {
	_, err := Client.Collection("users").Doc(userData.Email).Collection("books").Doc(docId).Delete(context.Background())
	return err
}
