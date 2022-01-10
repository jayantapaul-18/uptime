package routers

import (
	"fmt"
	"net/http"
	"text/template"
)

type Todo struct {
	Title string
	Done  bool
}
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
	Apis      []Api
}

type Api struct {
	Id       int
	Route    string
	Status   string
	Comments string
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.page.html")
}

func Admin(w http.ResponseWriter, r *http.Request) {
	parseTemplate, _ := template.ParseFiles("./templates/admin.page.html")
	data := TodoPageData{
		PageTitle: "Admin Page",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
		Apis: []Api{
			{Id: 1, Route: "/admin", Status: "ok", Comments: "healthy"},
			{Id: 2, Route: "/home", Status: "down", Comments: "not healthy"},
			{Id: 3, Route: "/ping", Status: "down", Comments: "not healthy"},
			{Id: 4, Route: "/healthcheck", Status: "down", Comments: "not healthy"},
		},
	}
	err := parseTemplate.Execute(w, data)
	if err != nil {
		fmt.Println("Error parsing html template:", err)
	}
}

// Render HTML
func renderTemplate(w http.ResponseWriter, html string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing html template:", err)
	}
}

// Render TMPL
func renderTMPLTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("../templates/" + tmpl)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing tmpl template:", err)
	}
}
