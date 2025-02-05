package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rioba-Ian/github-user-activity/models"
)

func DecodeResponse(data []byte) (*string, error) {
	var base models.Event

	if err := json.Unmarshal(data, &base); err != nil {
		return nil, fmt.Errorf("failed to decode data to json: %w", err)
	}

	switch base.Type {
	case "CreateEvent":
		var payload models.CreateEventPayload
		if err := json.Unmarshal(base.Payload, &payload); err != nil {
			return nil, fmt.Errorf("failed to decode create event payload data to json: %w", err)

		}

		activityMessage := ActivityResponse("created a new repository")

		return &activityMessage, nil

	case "WatchEvent":
		var payload models.WatchEventPayload
		if err := json.Unmarshal(base.Payload, &payload); err != nil {
			return nil, fmt.Errorf("failed to decode watch event payload data to json: %w", err)

		}

		activityMessage := ActivityResponse("Starred a new repo")

		return &activityMessage, nil

	case "PushEvent":
		var payload models.PushEventPayload
		if err := json.Unmarshal(base.Payload, &payload); err != nil {
			return nil, fmt.Errorf("failed to decode push event payload data to json: %w", err)

		}

		commits := len(payload.Commits)

		activityMessage := ActivityResponse(fmt.Sprintf("Pushed %d commit(s) to %s", commits, base.Repo.Name))

		return &activityMessage, nil

	case "PullRequestEvent":
		var payload models.PullRequestEventPayload
		if err := json.Unmarshal(base.Payload, &payload); err != nil {
			return nil, fmt.Errorf("failed to decode pull request event payload data to json: %w", err)

		}

		pullRequestAction := PullRequestHandler(payload.Action)

		activityMessage := ActivityResponse(pullRequestAction)

		return &activityMessage, nil

	case "IssueCommentEvent":
		var payload models.IssueCommentEventPayload
		if err := json.Unmarshal(base.Payload, &payload); err != nil {
			return nil, fmt.Errorf("failed to decode issue comment event payload data to json: %w", err)

		}

		var activityMessage string

		if payload.Action == "created" {
			activityMessage = "Created a new issue"
		} else if payload.Action == "Deleted" {
			activityMessage = "Deleted an issue"
		}

		return &activityMessage, nil

	default:
		return nil, fmt.Errorf("unknown type: %s", base.Type)
	}

}

func HandleResponse(username string) string {
	var events []models.Event
	fmt.Println("fetching data from gh...")
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))

	if err != nil {
		fmt.Printf("failed to get events %s\n", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		fmt.Printf("failed to decode body response from api %s", err)
	}

	latestEvent, _ := json.MarshalIndent(events[0], "", "	")

	fmt.Println("latest github user event", string(latestEvent))

	decodedResult, err := DecodeResponse(latestEvent)

	if err != nil {
		fmt.Printf("failed to decode response from events %s", err)
	}

	return *decodedResult

}

func ActivityResponse(message string) string {
	return message
}

func PullRequestHandler(action string) string {

	switch action {
	case "opened":
		return "opened new pull request"

	case "closed":
		return "closed pull request"

	case "review_requested":
		return "requested review on pull request"

	default:
		return "had a pull request activity"

	}

}
