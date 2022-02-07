package environmentvar

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

func GetEnv() {
	var nodeenv string = os.Getenv("ENV_NAME_GO_TEST")
	fmt.Printf("Reading Environment Variable: %s\n", nodeenv)
}

func SetEnv() {
	err := os.Setenv("NODE_ENV", "production")
	if err != nil {
		fmt.Println("Error when setting environment variable: ", err)
	}
	var nodeenv string = os.Getenv("NODE_ENV")
	fmt.Printf("Setting Environment Variable: %s\n", nodeenv)
}

func GetAllEnv() {
	fmt.Printf("Reading Environment Variable")
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%s: %s\n", pair[0], pair[1])

	}
}

type Envvariable struct {
	EnvName  string `json:"envName"`
	EnvValue string `json:"envValue"`
}

func CreateNewEnv(w http.ResponseWriter, r *http.Request) {
	color.Blue("Request to CreateNewEnv -----")
	var envvariable Envvariable
	json.NewDecoder(r.Body).Decode(&envvariable)
	log.Println("BODY : ", envvariable)
	fmt.Printf("type of a json.NewDecoder(r.Body).Decode is %T\n", envvariable)
	b, err2 := io.ReadAll(r.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	color.Blue("REQUEST from envvariable: ")
	newData, err := json.Marshal(envvariable)
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
	// Write to OS ENV variable
	enverr := os.Setenv(envvariable.EnvName, envvariable.EnvValue)
	fmt.Println(envvariable.EnvName, "    ", envvariable.EnvValue)
	if enverr != nil {
		fmt.Println("Error when setting environment variable: ", enverr)
	}
	// Add Response Header & Response Body
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(newData)

}
