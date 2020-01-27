package lib

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
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

// GO Templates
func ParseTemplates(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("%s", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "header", data)
	w.WriteHeader(http.StatusOK)

}

func MongoToString(seller interface{}) string {
	json, err := json.Marshal(seller)

	if err != nil {
		fmt.Printf("JSON Error: %v on %s", err, seller)
		return ""
	}

	return string(json)
}
