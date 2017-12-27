package main

import (
	//"fmt"
	"bytes"
	"github.com/gorilla/mux"
	//"html"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// run --> go get github.com/gorilla/mux
	// in terminal to install the external package for mux

	// We need to create a router
	rt := mux.NewRouter().StrictSlash(true)

	// Add the "index" or root path
	rt.HandleFunc("/", Index)

	// Fire up server
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}

// Index is the "index" handler
func Index(w http.ResponseWriter, r *http.Request) {
	// Fill out the page data for index
	pd := PageData{
		Title: "Index Page",
		Body:  "This is the body of the index page.",
	}

	// Render a template with our page data
	tmpl, err := render(pd)

	// if we get an error, write it out and exit
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// All went well, so write out the template
	w.Write([]byte(tmpl))

	//fmt.Fprintf(w, "Hello world from %q", html.EscapeString(r.URL.Path))
}

const html = `
<html>
	<head>
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{.Body}}
	</body>
</html>
`

func render(pd PageData) (string, error) {
	// Parse the template
	tmpl, err := template.New("index").Parse(html)
	if err != nil {
		return "", err
	}

	// We need somewhere to write the executed template to
	var out bytes.Buffer

	// Render the template with the data we passed in
	if err := tmpl.Execute(&out, pd); err != nil {
		// If we couldn't render, return an error
		return "", err
	}

	// Return the template
	return out.String(), nil
}

type PageData struct {
	Title string
	Body  string
}
