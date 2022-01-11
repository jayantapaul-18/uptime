package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"jayantapaul-18/uptime/pkg/config"
	"jayantapaul-18/uptime/pkg/custommiddleware"
	"jayantapaul-18/uptime/pkg/helpers"
	"jayantapaul-18/uptime/pkg/mypackage"
	"jayantapaul-18/uptime/pkg/mysqldb"
	"jayantapaul-18/uptime/pkg/routers"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	// _ "github.com/go-sql-driver/mysql"
	// "path/filepath"
	// c "./config"
)

const portNumber = ":3088"

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

type healthzResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

type Employee struct {
	Id   int
	Name string
	City string
}

func init() {
	// Concurrency || Parallelism
	fmt.Println("Total Number of CPU: ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Go Version: ", runtime.Version())
	time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	color.Blue("==========================\n")
	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline)
	// Create a custom print function for convenience
	red := color.New(color.FgRed).PrintfFunc()
	red("Red Error logs ")
	c.Println("Ending initilazation...")
}

func main() {

	log.Println("Go Server starting ... Time: ", time.Now())
	log.Println("OS:", os.Getenv("GOPATH"))
	microlevel := map[string]int{
		"LOW":    0,
		"MEDIUM": 5,
		"HIGH":   10,
	}
	fmt.Println(microlevel)
	printMap(microlevel)
	var flagNoColor = flag.Bool("no-color", false, "Disable color output")
	if *flagNoColor {
		color.NoColor = true // disables colorized output
	}
	// Production mode Config flag
	inProduction := flag.Bool("production", true, "Application is in production mode")
	app.InProduction = *inProduction
	log.Println("InProduction: ", app.InProduction)
	flag.Parse()
	if !*inProduction || *inProduction == false {
		fmt.Println("Missing InProduction configuration or sets to false")
		os.Exit(1)
	}

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog
	helpers.NewHelpers(&app)
	// CHI - Server // https://go-chi.io/
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	r.Use(middleware.CleanPath)
	r.Use(middleware.GetHead)
	r.Use(middleware.Compress(5, "text/html", "text/css"))
	r.Use(middleware.Timeout(time.Second * 10))
	// Custom middleware [panic: chi: all middlewares must be defined before routes on a mux]
	// r.Use(custommiddleware.NoSurf)
	//r.Use(custommiddleware.WriteToConsoleLog)
	r.Use(custommiddleware.DumpRequest)
	// r.Use(custommiddleware.RequestID)
	// r.Use(custommiddleware.RequestTime)
	// r.Use(custommiddleware.WhiteList)
	// r.Use(custommiddleware.AddHeaders)
	r.Mount("/app/x/debug", middleware.Profiler())

	mypackage.Hello() // custom package
	// DNS Check
	//dnsrun.DnsCheck()
	// DB Connection
	db := mysqldb.DBConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}
	defer db.Close()
	log.Println(res)

	r.Post("/app/v1/create-profile", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
		w.Write([]byte("Signup Done"))
	})

	r.Post("/app/v1/signup", mysqldb.CreateNewArticle)

	r.Post("/app/v1/login", mysqldb.Login)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up"))
	})

	// r.Get("/app/v1/db-search", mysqldb.DBSearch)

	r.Get("/app/v1/healthz", func(w http.ResponseWriter, r *http.Request) {
		resp := healthzResponse{
			OK:      true,
			Message: "Healthy",
		}

		out, err := json.MarshalIndent(resp, "", "      ")
		if err != nil {
			log.Println("JSON MarshalIndent error:", err)
		}
		if govalidator.IsJSON(string(out)) == true {
			log.Println("Valid JSON")
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	})
	r.Get("/app/v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		client := http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get("http://localhost:3088/app/v1/healthz")
		// err = errors.New("New Error !")
		if err != nil {
			helpers.ServerError(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	})

	r.Get("/home", routers.Home)
	r.Get("/index", routers.Index)
	r.Get("/admin", routers.Admin)
	r.Get("/alpine", routers.Alpine)
	r.Get("/profile", routers.Profile)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	color.Magenta("Server Continue Running ...")
	color.Cyan("##  Server  http://localhost:3088/ping ##")
	addr := getServiceAddress()
	log.Printf("Server running on http://%s", addr)
	// http.ListenAndServe(portNumber, r)
	if err := http.ListenAndServe(portNumber, r); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Server END 	===================""
// Get server address
func getServiceAddress() string {
	if env := os.Getenv("PORT_BEHIND_PROXY"); env != "" {
		return ":" + env
	}
	if env := os.Getenv("VIRTUAL_PORT"); env != "" {
		return ":" + env
	}

	return portNumber
}

func printMap(m map[string]int) {
	for level, val := range m {
		fmt.Println("Micro Level value for ", level, "is", val)
	}
}
