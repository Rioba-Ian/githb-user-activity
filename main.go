package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Event struct {
	Id      string  `json:"id"`
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	RespositoryId int64    `json:"repository_id"`
	PushId        int64    `json:"push_id"`
	Size          int64    `json:"size"`
	DistinctSize  int64    `json:"distinct_size"`
	Ref           string   `json:"ref"`
	Head          string   `json:"head"`
	Before        string   `json:"before"`
	Commits       []Commit `json:"commits"`
}

type Commit struct {
	Sha      string `json:"sha"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	Url      string `json:"url"`
}

func main() {
	fmt.Println("fetching data from gh.")

	latestEvent := getEvents("Rioba-Ian")

	_ = latestEvent
}

func getEvents(username string) Event {
	fmt.Println("The username is", username)
	_ = username
	var events []Event

	resp, err := http.Get("https://api.github.com/users/kamranahmedse/events")

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Failed to get events %s\n", err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&events)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("latest event", events[0])

	return events[0]
}
