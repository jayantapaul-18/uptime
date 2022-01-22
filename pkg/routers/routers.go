package routers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"text/template"

	"github.com/gookit/config"
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

func Alpine(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "alpine.page.html")
}
func Profile(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "profile.page.html")
}

type DNSResponse struct {
	DNS       string `json:"dns"`
	IPADDRESS []net.IP
	CNAME     string `json:"CNAME"`
}

func DCheck(w http.ResponseWriter, r *http.Request) {
	// renderTemplate(w, "alpine.page.html")
	DNS_URL_CHECK, _ := config.String("DNS_URL_CHECK")
	fmt.Print("DNS_URL_CHECK:", DNS_URL_CHECK)

	cname, _ := net.LookupCNAME(DNS_URL_CHECK)
	fmt.Println(" \ncname:", cname)
	iprecords, _ := net.LookupIP(DNS_URL_CHECK)
	fmt.Printf("%T\t", iprecords)
	n := len(iprecords)
	addrs := make([]net.TCPAddr, 0, n)
	for i := 0; i < n; i++ {
		ip := iprecords[i]
		addrs = append(addrs, net.TCPAddr{
			IP: ip,
		})
	}

	resp := DNSResponse{
		DNS:       DNS_URL_CHECK,
		IPADDRESS: iprecords,
		CNAME:     cname,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
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
