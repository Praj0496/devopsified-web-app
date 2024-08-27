// Test the main function

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStaticFiles(t *testing.T) {
	// Test serving the index.html file
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.FileServer(http.Dir("./static"))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Verify the content type for HTML
	expected := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, expected)
	}
}

func TestStaticCSSFile(t *testing.T) {
	// Test serving a CSS file
	req, err := http.NewRequest("GET", "/css/style.css", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.FileServer(http.Dir("./static"))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Verify the content type for CSS
	expected := "text/css; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, expected)
	}
}

func TestStaticImageFile(t *testing.T) {
	// Test serving an image file
	req, err := http.NewRequest("GET", "/images/logo.png", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.FileServer(http.Dir("./static"))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Verify the content type for PNG
	expected := "image/png"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, expected)
	}
}

func TestStaticJSFile(t *testing.T) {
	// Test serving the JS file located at /js/site.js
	req, err := http.NewRequest("GET", "/js/site.js", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.FileServer(http.Dir("./static"))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Accept both 'text/javascript' and 'application/javascript' as valid content types
	expectedContentTypes := []string{
		"text/javascript; charset=utf-8",
		"application/javascript; charset=utf-8",
	}

	contentType := rr.Header().Get("Content-Type")
	valid := false
	for _, expected := range expectedContentTypes {
		if contentType == expected {
			valid = true
			break
		}
	}

	if !valid {
		t.Errorf("handler returned unexpected content type: got %v want %v or %v",
			contentType, expectedContentTypes[0], expectedContentTypes[1])
	}
}
