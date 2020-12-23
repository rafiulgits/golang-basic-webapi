package main

import "net/http"

func main() {
	//inline handler function
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"message" : "hello world"}`))
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
