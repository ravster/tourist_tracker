package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// NEXT:
// Write entries and exits into DB using channel-receivers
// Script to hammer these 2 endpoints
// Webpage to show some stats

type NewEntry struct {
	passportNumber string
	portOfEntryId int
}
var newEntryChannel chan NewEntry
type NewExit struct {
	passportNumber string
	portOfExitId int
}
var newExitChannel chan NewExit

// curl -v localhost:3001/entryNew -d passportNumber=yyh665 -d portId=23
func entryNew(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO add user-id to this when we eventually get past basic functionality and want to
	// add authentication. This will be used to send an event to the right user at the port of entry.
	passportNumber := r.PostFormValue("passportNumber")
	portId := r.PostFormValue("portId")
	if passportNumber == "" {
		http.Error(w, "Error: passportNumber is blank", http.StatusUnprocessableEntity)
		return
	}
	portOfEntryId, err := strconv.Atoi(portId)
	if err != nil {
		http.Error(w, "portId should be int", http.StatusUnprocessableEntity)
		return
	}

	newEntry := NewEntry{ passportNumber: passportNumber, portOfEntryId: portOfEntryId }
	// Prodnote: In a real system this would go into a proper background worker system, backed by
	// something like SQS (a data-residency safe version of it). This prevents issues with
	// this particular server going down due to a power outage or network failure, and then this
	// entry not actually getting recorded.
	newEntryChannel <- newEntry

	fmt.Printf("enh. pn=%s. port=%s\n", passportNumber, portId)
	w.WriteHeader(http.StatusAccepted)
}

// curl -v localhost:3001/exitNew -d passportNumber=yyh665 -d portId=23
func exitNew(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// TODO can this be parsed straight into a struct?
	passportNumber := r.PostFormValue("passportNumber")
	portId := r.PostFormValue("portId")
	if passportNumber == "" {
		http.Error(w, "Error: passportNumber is blank", http.StatusUnprocessableEntity)
		return
	}
	portOfExitId, err := strconv.Atoi(portId)
	if err != nil {
		http.Error(w, "portId should be int", http.StatusUnprocessableEntity)
		return
	}

	newExit := NewExit{ passportNumber: passportNumber, portOfExitId: portOfExitId }
	// Prodnote: In a real system this would go into a proper background worker system, backed by
	// something like SQS (a data-residency safe version of it). This prevents issues with
	// this particular server going down due to a power outage or network failure, and then this
	// exit not actually getting recorded.
	newExitChannel <- newExit

	fmt.Printf("exitNew. pn=%s. port=%s\n", passportNumber, portId)
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	fmt.Println("Started program")

	// Setup global channels.
	// Lets say this system is used in a small nation with 2-3 internations ports. Assuming 10 entry
	// immigration lines at each port, we'd have a max of 30 people entering the country at the same
	// time. A buffered channel of size 500 should handle all of that fine.
	newEntryChannel = make(chan NewEntry, 500) // TODO receiver
	newExitChannel = make(chan NewExit, 500) // TODO receiver

	http.HandleFunc("/entryNew", entryNew) // POST
	http.HandleFunc("/exitNew", exitNew) // POST
	fmt.Println("Starting server at :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
