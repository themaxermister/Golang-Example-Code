package lib

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)
		log.Println(req)
		next.ServeHTTP(w, r)
		log.Println(req, "completed in", time.Now().Sub(start))
	})
}

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b := struct {
			Title        template.HTML
			BusinessName string
			Slogan       string
		}{
			Title:        template.HTML("Business &verbar; Landing"),
			BusinessName: "Business,",
			Slogan:       "we get things done.",
		}
		err := templates.ExecuteTemplate(w, "base", &b)
		if err != nil {
			http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func TestIndex(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("TestIndex: couldn't create HTTP GET request: %v", err)
	}

	rec := httptest.NewRecorder()

	Index().ServeHTTP(rec, req)

	res := rec.Result()
	defer func() {
		err := res.Body.Close()
		if err != nil {
			t.Fatalf("TestIndex: couldn't close response body: %v", err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		t.Errorf("TestIndex: got status code %v, but want: %v", res.StatusCode, http.StatusOK)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("TestIndex: could not read response body: %v", err)
	}

	if len(string(body)) == 0 {
		t.Errorf("TestIndex: unexpected empty response body")
	}
}

func Public() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
}
