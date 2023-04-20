package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bonhokage06/lenslocked/controllers"
	"github.com/bonhokage06/lenslocked/helpers"
	middlewares "github.com/bonhokage06/lenslocked/middleware"
)

func TestHomeIndexRequest(t *testing.T) {
	helpers.CurrentWorkingDirectory()
	resp, err := http.Get("http://localhost:3000")
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	dataString := string(data)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}
	if !strings.Contains(dataString, "Welcome to my awesome site!") {
		t.Error("Expected body to contains 'Welcome to my awesome site!'")
	}
}
func TestContactIndexRequest(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/contact")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	dataString := string(data)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}
	if !strings.Contains(dataString, "Welcome to Contact!") {
		t.Error("Expected body to contains 'Welcome to Contact!'")
	}
}
func TestFaqIndexRequest(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/faq")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	dataString := string(data)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}
	if !strings.Contains(dataString, "Welcome to FAQ!") {
		t.Error("Expected body to contains 'Welcome to FAQ!'")
	}
}
func TestStaticRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/static/js/main.js", nil)
	w := httptest.NewRecorder()
	middlewares.StaticHandler(controllers.Static()).ServeHTTP(w, req)
	resp := w.Result()

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	dataString := string(data)
	fmt.Println(dataString)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}
	if len(dataString) == 0 {
		t.Error("Expected body to contains 'main.js' content")
	}
}

func TestSignUpRequest(t *testing.T) {
	helpers.CurrentWorkingDirectory()
	resp, err := http.Get("http://localhost:3000/signup")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	dataString := string(data)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}
	if !strings.Contains(dataString, "Welcome to Signup!") {
		t.Error("Expected body to contains 'Welcome to Signup!'")
	}
}
