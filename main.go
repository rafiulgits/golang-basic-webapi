package main

import (
	"net/http"
)

const (
	HttpGET    = "GET"
	HttpPOST   = "POST"
	HttpPUT    = "PUT"
	HttpDELETE = "DELETE"
)

func main() {
	//inline handler function
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		switch req.Method {
		case HttpGET:
			w.WriteHeader(200)
			w.Write([]byte(`{"message" : "data fetch"}`))
		case HttpPOST:
			w.WriteHeader(201)
			w.Write([]byte(`{"message" : "data created"}`))
		case HttpPUT:
			w.WriteHeader(200)
			w.Write([]byte(`{"message" : "data updated"}`))
		case HttpDELETE:
			w.WriteHeader(204)
		default:
			w.WriteHeader(200)
			w.Write([]byte(`{"message" : "no methods matched"}`))
		}

	})

	//passing handler function
	http.HandleFunc("/greet", worldHandler)

	// :8090 meaning serve all hosts available in local machine
	// or simply 0.0.0.0:8090
	http.ListenAndServe(":8090", nil)
}

func worldHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"message" : "greet"}`))
}
