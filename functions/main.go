package main

import (
	"bufio"
	"html/template"
	"net/http"
	"os"
	"time"
)

var (
	port = ":8002"
)

// Log - is the structure of one log
type Log struct {
	Title       string   `json:"title"`
	Body        string   `json:"body"`
	Tags        []string `json:"tags"`
	LastUpdated time.Time
}

func main() {

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)
	http.HandleFunc("/home/", HomeHandler)

	//browser.OpenURL("localhost" + port)
	http.ListenAndServe(port, nil)
}

func scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	return line
}

func handler(w http.ResponseWriter, r *http.Request) {

	//the template file to render, usually as html, shows static data and also data added from next part
	tmpl := template.Must(template.ParseFiles("layout.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)

		return
	}

	//the information taken feom the webpage in the form of a LOG
	details := Log{
		Title:       r.FormValue("title"),
		Body:        r.FormValue("body"),
		Tags:        getTags(r.FormValue("body")),
		LastUpdated: time.Now(),
	}

	//create the log from the detail. the functions connects to mlab.
	details.CreateLog()

	//tmpl.Execute(w, data)
	tmpl.Execute(w, struct{ Success bool }{true})

}

// HomeHandler - takes care of the data going to the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("layout.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)

		return
	}

	details := Log{
		Title:       r.FormValue("title"),
		Body:        r.FormValue("body"),
		Tags:        getTags(r.FormValue("body")),
		LastUpdated: time.Now(),
	}

	details.CreateLog()

	//tmpl.Execute(w, data)
	tmpl.Execute(w, struct{ Success bool }{true})

}
