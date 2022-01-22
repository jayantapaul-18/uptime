package helpers

import (
	"fmt"
	"jayantapaul-18/uptime/pkg/localconfig"
	"net/http"
	"runtime/debug"
)

var app *localconfig.AppConfig

// Sets up config for helpers
func NewHelpers(a *localconfig.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of ", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	app.ErrorLog.Println("Server error with status of ", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
