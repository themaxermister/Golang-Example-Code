package lib

import(
	"html/template"
)

// Variables
var Tpl = template.Must(template.ParseFiles("index.html")) // Templates
var ApiKey *string
