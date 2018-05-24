
## save file

```go
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
```


## part 1
```go
package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

// Log - is the structure of one log
type Log struct {
	Title       string   `json:"title"`
	Body        string   `json:"body"`
	Tags        []string `json:"tags"`
	LastUpdated time.Time
}

// MainPage
type MainPage struct {
	Logs []Log
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/*data := MainPage{
			Logs:    GetLogs(),
			Success: false,
		}
		*/
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

	})

	http.ListenAndServe(":8002", nil)
}

// MakeLog - Saves a log to DB with only title nad text input
func MakeLog(title, body string) {

	l := Log{
		Title:       title,
		Body:        body,
		Tags:        getTags(body),
		LastUpdated: time.Now(),
	}
	l.CreateLog()
}

func scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	return line
}

// this works on the command line
func writeCommandLineLog() {
	fmt.Print("What is the title?\n")
	title := scanner()
	fmt.Println("Type the log\n ")
	body := scanner()

	MakeLog(title, body)
	//body := scanner()
	//MakeLog(title, "#spank this is a # foshizzle my guchi is #cashizzle")

	CallClear()
	fmt.Printf("The log is saved with title: %s", title)
}

```

<form method="POST">
			<label>Title:</label>
			<br />
			<input type="text" name="title">
			<br />
			<label>Body:</label>
			<br />
			<input type="text" name="body">
			<br />
			<input type="submit">
		</form>