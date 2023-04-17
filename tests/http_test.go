package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bonhokage06/lenslocked/controllers"
	"github.com/bonhokage06/lenslocked/helpers"
)

func TestHomeIndexRequest(t *testing.T) {
	helpers.CurrentWorkingDirectory()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	w := httptest.NewRecorder()
	controllers.HtmlHandler(controllers.Index(nil, "home.gohtml", "partials/layout-parts.gohtml")).ServeHTTP(w, req)
	resp := w.Result()

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
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
	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/contact", nil)
	w := httptest.NewRecorder()
	controllers.HtmlHandler(controllers.Index(nil, "contact.gohtml", "partials/layout-parts.gohtml")).ServeHTTP(w, req)
	resp := w.Result()

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
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
	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/faq", nil)
	w := httptest.NewRecorder()
	controllers.HtmlHandler(controllers.Index(nil, "faq.gohtml", "partials/layout-parts.gohtml")).ServeHTTP(w, req)
	resp := w.Result()

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
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
	controllers.StaticHandler(controllers.IndexStatic()).ServeHTTP(w, req)
	resp := w.Result()

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
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
