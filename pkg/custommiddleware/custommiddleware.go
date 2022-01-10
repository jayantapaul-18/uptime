package custommiddleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/justinas/nosurf"
)

func WriteToConsoleLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request -->")
		next.ServeHTTP(w, r)
	})
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r2 := new(http.Request)
		*r2 = *r
		uuidn := uuid.New()
		fmt.Println(uuidn)
		// r2.Header.Set("X-Request-Id", uuidn)
		// w.Header().Add("X-Request-Id", string(uuidn))
		next.ServeHTTP(w, r2)
	})
}

func RequestTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "requestTime", time.Now().Format(time.RFC3339))
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

type MyResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func (mrw *MyResponseWriter) Write(p []byte) (int, error) {
	return mrw.buf.Write(p)
}
func DumpRequest(next http.Handler) http.Handler {
	// r.Body stream can be copy only once - multiple copy solution below
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// requestDump, err := httputil.DumpRequest(r, true)
		// if err != nil {
		// 	log.Println(err.Error())
		// } else {
		// 	color.Blue("REQUEST - DUMP STARTED ===================")
		// 	color.Green("requestDump: ", string(requestDump))
		// 	color.Blue("REQUEST - DUMP END     ===================")
		b, err2 := io.ReadAll(r.Body)
		if err2 != nil {
			log.Fatal(err2)
		}
		color.Blue("REQUEST - BODY STARTED ===================")
		color.Yellow(string(b))
		color.Blue("REQUEST - BODY END ===================")

		// And now set a new body, which will simulate the same data we read:
		r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		// Create a response wrapper:
		mrw := &MyResponseWriter{
			ResponseWriter: w,
			buf:            &bytes.Buffer{},
		}
		// Call next handler, passing the response wrapper:
		next.ServeHTTP(mrw, r)
		// Now inspect response, and finally send it out:
		// (You can also modify it before sending it out!)
		if _, err := io.Copy(w, mrw.buf); err != nil {
			log.Printf("Failed to send out response: %v", err)
		}
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func WhiteList(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("WhiteList ..")
		log.Println("Whitelisted route for Admin URL: ", r.URL.Path)
		if r.URL.Path == "/app/v1/admin" {
			return
		}
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
		log.Println("Whitelisted route for Admin URL: ", r.URL.Path)
	})
}

func AddHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
