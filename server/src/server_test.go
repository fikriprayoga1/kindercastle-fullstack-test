package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var bookJSON = "{\"title\":\"Berjaya\",\"page_total\":203,\"written_by\":\"Joko Widodo\"}\n"

func TestCreateBook(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/book", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{
		saveBookData: func(u *Book) error {
			return nil
		},
	}

	// Assertions
	if assert.NoError(t, h.createBookHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, bookJSON, rec.Body.String())
	}
}

func TestReadUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{
		readBookData: func() ([]Book, error) {
			return []Book{}, nil
		},
	}

	// Assertions
	if assert.NoError(t, h.readBooksHandler(c)) {
		assert.Equal(t, http.StatusFound, rec.Code)
		assert.Equal(t, "[]\n", rec.Body.String())
	}
}

func TestUpdateUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/book/usdkafgkuyasd", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{
		updateBookData: func(docId string, u *Book) error {
			return nil
		},
	}

	// Assertions
	if assert.NoError(t, h.updateBookHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, bookJSON, rec.Body.String())
	}
}

func TestDeleteUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/book/usdkafgkuyasd", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{
		deleteBookData: func(docId string) error {
			return nil
		},
	}

	// Assertions
	if assert.NoError(t, h.deleteBookHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Book Deleted", rec.Body.String())
	}
}
