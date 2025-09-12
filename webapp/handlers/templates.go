package handlers

import "text/template"

var Templates = template.Must(template.ParseGlob("templates/*.tmpl"))
