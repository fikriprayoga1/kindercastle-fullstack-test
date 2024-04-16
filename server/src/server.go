package main

import (
	"context"
	"log"
	"net/http"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/option"
)

const port string = "1323"

var (
	app      *firebase.App
	userData *auth.UserRecord
	Client   *firestore.Client
)

type (
	Book struct {
		DocId     string `json:"doc_id,omitempty"`
		Title     string `json:"title,omitempty"`
		PageTotal int    `json:"page_total,omitempty"`
		WrittenBy string `json:"written_by,omitempty"`
	}

	Handler struct {
		saveBookData   func(u *Book) error
		readBookData   func() ([]Book, error)
		updateBookData func(docId string, u *Book) error
		deleteBookData func(docId string) error
	}
)

func main() {
	initFirebase()
	initServer()
	defer Client.Close()

}

func bearerTokenHandler(uid string, c echo.Context) (bool, error) {
	// Get an auth client from the firebase.App
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	userData, err = client.GetUser(context.Background(), uid)
	if err != nil {
		log.Printf("error getting user %s: %v\n", uid, err)
		return false, nil
	}
	log.Printf("Successfully fetched user data: %v\n", userData.Email)

	return true, nil
}

func initFirebase() {
	var err error
	opt := option.WithCredentialsFile("./kindercastle-private-key.json")
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	Client, err = app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

}

func initServer() {
	log.Printf("Server listener started.\n\n")

	e := echo.New()
	e.Use(middleware.KeyAuth(bearerTokenHandler))

	h := Handler{
		saveBookData:   SaveBookData,
		readBookData:   ReadBookData,
		updateBookData: UpdateBookData,
		deleteBookData: DeleteBookData,
	}

	e.POST("/book", h.createBookHandler)
	e.GET("/books", h.readBooksHandler)
	e.PUT("/book/:doc_id", h.updateBookHandler)
	e.DELETE("/book/:doc_id", h.deleteBookHandler)
	e.Logger.Fatal(e.Start(":" + port))
}

func (h *Handler) createBookHandler(c echo.Context) error {
	u := new(Book)
	if err := c.Bind(u); err != nil {
		return err
	}

	err := h.saveBookData(u)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)

		return c.String(http.StatusInternalServerError, "Error Server")
	}

	return c.JSON(http.StatusCreated, u)

}

func (h *Handler) readBooksHandler(c echo.Context) error {
	result, err := h.readBookData()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error Server")
	}

	return c.JSON(http.StatusFound, result)
}

func (h *Handler) updateBookHandler(c echo.Context) error {
	docId := c.Param("doc_id")

	u := new(Book)
	if err := c.Bind(u); err != nil {
		return err
	}

	err := h.updateBookData(docId, u)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return c.String(http.StatusInternalServerError, "Error Server")
	}

	return c.JSON(http.StatusOK, u)
}

func (h *Handler) deleteBookHandler(c echo.Context) error {
	docId := c.Param("doc_id")

	err := h.deleteBookData(docId)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return c.String(http.StatusInternalServerError, "Error Server")
	}

	return c.String(http.StatusOK, "Book Deleted")
}
