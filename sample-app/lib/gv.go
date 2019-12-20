package lib

import "html/template"

var templates = template.Must(template.ParseFiles("./templates/base.html", "./templates/body.html"))
