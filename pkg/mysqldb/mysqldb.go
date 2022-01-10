package mysqldb

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
)

// You should make your function exportable with an uppercase like func DBConn()
func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "dbuser"
	dbPass := "dbuser"
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		log.Println("Failed to connect to database: ", err.Error())
	}
	log.Println("db connection:", db)
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func DBSearch() (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("db search"))
	return
}

func createUsersTable(db *sql.DB) {
	st, err0 := db.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER NOT NULL AUTO_INCREMENT,userName varchar(255),password varchar(20),urlNum int,PRIMARY KEY (id))")
	if err0 != nil {
		panic(err0.Error())
	}
	_, err1 := st.Exec()
	if err1 != nil {
		panic(err1.Error())
	}

}

func Insert(w http.ResponseWriter, r *http.Request) {
	log.Println("Insert Request:")
	db := DBConn()
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		log.Println(body)
		// requestDump, err := httputil.DumpRequest(r, true)
		// if err != nil {
		// 	log.Println(err.Error())
		// } else {
		// 	log.Println(string(requestDump))

		// 	b, err2 := io.ReadAll(r.Body)
		// 	if err2 != nil {
		// 		log.Fatal(err2)
		// 	}
		// 	log.Println(string(b))
		// }
		// name := r.FormValue("name")
		// city := r.FormValue("city")
		// log.Println("Request:", name)
		//insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(name,city)")
		// if err != nil {
		// 	panic(err.Error())
		// }
		// insForm.Exec(name, city)
		// log.Println("INSERT: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	w.WriteHeader(201)
	w.Write([]byte("Signup Done"))
	// http.Redirect(w, r, "/", 301)
}

type MyResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func (mrw *MyResponseWriter) Write(p []byte) (int, error) {
	return mrw.buf.Write(p)
}

type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
	Summary string `json:"Summary"`
}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	color.Blue("Request to CreateNewArticle -----")
	var article Article
	json.NewDecoder(r.Body).Decode(&article)
	log.Println("BODY : ", article)
	fmt.Printf("type of a json.NewDecoder(r.Body).Decode is %T\n", article)
	b, err2 := io.ReadAll(r.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	color.Blue("REQUEST from article: ")
	newData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("type of a json.Marshal is %T\n", newData)
		fmt.Println(string(newData))
	}
	out, err := json.MarshalIndent(newData, "", "      ")
	if err != nil {
		log.Println("JSON MarshalIndent error:", err)
	}
	fmt.Printf("type of a json.MarshalIndent is %T\n", out)
	color.Red(string(out))
	color.Yellow(string(b))
	fmt.Printf("type of b = io.ReadAll(r.Body) is %T\n", b)
	// Write to Database
	db := DBConn()
	sqlStatement := `INSERT IGNORE INTO article(Id,Title,Content,Summary) VALUES (?, ?, ?, ?)`
	log.Println(sqlStatement)
	_, err = db.Exec(sqlStatement, article.Id, article.Title, article.Content, article.Summary)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Add Response Header & Response Body
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(newData)

}

func Loginmw(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		// Work / inspect body. You may even modify it!

		// And now set a new body, which will simulate the same data we read:
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// Create a response wrapper:
		mrw := &MyResponseWriter{
			ResponseWriter: w,
			buf:            &bytes.Buffer{},
		}

		// Call next handler, passing the response wrapper:
		handler.ServeHTTP(mrw, r)

		// Now inspect response, and finally send it out:
		// (You can also modify it before sending it out!)
		if _, err := io.Copy(w, mrw.buf); err != nil {
			log.Printf("Failed to send out response: %v", err)
		}
	})
}
