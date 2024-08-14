package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func (t Info) GetAffiliationDetailInfo() string {
	return "Have 31 Divisions"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name:    "Breuce Wayne",
			Gender:  "male",
			Hobbies: []string{"Reading Books", "Traveling", "Shopping"},
			Info:    Info{"Wayne Enterprise", "Gotham City"},
		}

		tmpl := template.Must(template.ParseFiles("view.html"))

		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("server started on localhost:9000")
	http.ListenAndServe(":9000", nil)
}
