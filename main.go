package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! You requested: %s\n", r.URL.Path)
}

func entryNew(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// TODO can this be parsed straight into a struct?
	passportNumber := r.PostFormValue("passportNumber")
	portId := r.PostFormValue("portId")
	// TODO Shove this into a channel. The reader of the chan should know where to write the data.
	// Reader should also (eventually) check this against the ban-list.
	fmt.Printf("enh. pn=%s. port=%s\n", passportNumber, portId)
	w.WriteHeader(http.StatusAccepted)
}

func exitNew(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// TODO can this be parsed straight into a struct?
	passportNumber := r.PostFormValue("passportNumber")
	portId := r.PostFormValue("portId")
	// TODO Shove this into a channel. The reader of the chan should know where to write the data.
	fmt.Printf("exitNew. pn=%s. port=%s\n", passportNumber, portId)
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	fmt.Println("Started program")

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/entryNew", entryNew) // POST
	http.HandleFunc("/exitNew", exitNew) // POST
	fmt.Println("Starting server at :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
