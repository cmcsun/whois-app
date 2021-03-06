package main

import (
	"encoding/json"
	"github.com/domainr/whois"
	"io"
	"log"
	"net/http"
)

func main() {
	// Create a fileServer handler that serves our static files.
	fileServer := http.FileServer(http.Dir("static/"))

	// Tell the http library how we want to handle requests.
	// For now, we simply pass the request to the fileServer.
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fileServer.ServeHTTP(w, r)
		},
	)

	// Handle the POST request on /whois
	// The client will send a url-encoded request like:
	//     data=8.8.8.8
	http.HandleFunc("/whois", func(w http.ResponseWriter, r *http.Request) {
		// Verify this is POST (not e.g. GET or DELETE).
		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Extract the encoded data to perform the whois on.
		data := r.PostFormValue("data")

		// Perform the whois query.
		result, err := whoisQuery(data)

		// Return a JSON-encoded response.
		if err != nil {
			jsonResponse(w, Response{Error: err.Error()})
			return
		}
		jsonResponse(w, Response{Result: result})
	})

	// Finally, start the HTTP server
	// If anything goes wrong, the log.Fatal call will output
	// the error to the console and exit the application.
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

type Response struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

func whoisQuery(data string) (string, error) {
	// Run whois on the user-specified object.
	response, err := whois.Fetch(data)
	if err != nil {
		return "", err
	}
	return string(response.Body), nil
}

func jsonResponse(w http.ResponseWriter, x interface{}) {
	// JSON-encode x.
	bytes, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	// Write the encoded data to the ResponseWriter.
	// This will send the response to the client.
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	//simple health check
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

/*
In progress:
Use Middleware for mux router instead of constantly needing to call
http.ResponseWriter Content-Type header globally for all
API functions that I have


func commonMiddleWare(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http Request){
		w.Header().Add("Content-Type","application/json")
		next.ServeHTTP(w,r)
	}
}

*/

//func
