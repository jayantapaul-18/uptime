package environmentvar

import (
	"fmt"
	"os"
)

func GetEnv() {
	var nodeenv string = os.Getenv("NODE_ENV")
	fmt.Printf("Reading Environment Variable NODE_ENV: %s\n", nodeenv)
}

func SetEnv() {
	err := os.Setenv("NODE_ENV", "production")
	if err != nil {
		fmt.Println("Error when setting environment variable: ", err)
	}
	var nodeenv string = os.Getenv("NODE_ENV")
	fmt.Printf("Setting Environment Variable NODE_ENV: %s\n", nodeenv)
}
