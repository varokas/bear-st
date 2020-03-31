package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/ratelimit"
)

var rl = ratelimit.New(100) // per second

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/captcha", captchaHandler)

	port := "3000"
	fmt.Println("Server Started at: " + port)
	http.ListenAndServe(":"+port, r)
}

type Message struct {
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	TotalNs   int64  `json:"totalNs"`
}

type Error struct {
	Error string `json:"error"`
}

const dateFormat = "2006-01-02T15:04:05.999999"

func createMessage(name string, startTime time.Time, endTime time.Time) Message {
	return Message{name, startTime.Format(dateFormat), endTime.Format(dateFormat), endTime.Sub(startTime).Nanoseconds()}
}

func captchaHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// Do stuff
	rl.Take()

	end := time.Now()

	writeResponse(w, "home", start, end)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// Do Nothing
	end := time.Now()

	writeResponse(w, "home", start, end)
}

func writeResponse(w http.ResponseWriter, endpoint string, start time.Time, end time.Time) {
	o, err := json.Marshal(createMessage(endpoint, start, end))
	if err == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(o))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		errJSON, _ := json.Marshal(Error{err.Error()})
		fmt.Fprintf(w, string(errJSON))
	}
}
