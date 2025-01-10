package models

import "encoding/json"

type Event struct {
	Id      string          `json:"id"`
	Type    string          `json:"type"`
	Repo    Repo            `json:"repo"`
	Payload json.RawMessage `json:"payload"`
}

type Repo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type PushEventPayload struct {
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

type PullRequestEventPayload struct {
	Action string `json:"action"`
	Number int64  `json:"number"`
	PullRequest
}

type PullRequest struct {
	Url     string `json:"url"`
	Id      int64  `json:"id"`
	NodeId  string `json:"node_id"`
	HtmlUrl string `json:"html_url"`
	Title   string `json:"title"`
	State   string `json:"state"`
}

type IssueCommentEventPayload struct {
	Action  string  `json:"action"`
	Issue   Issue   `json:"url"`
	Comment Comment `json:"comment"`
}

type Issue struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

type Comment struct {
	Url  string `json:"url"`
	Body string `json:"body"`
}

// stars a repository
type WatchEventPayload struct {
	Action string `json:"action"`
}

type CreateEventPayload struct {
	Ref          string `json:"ref"`
	RefType      string `json:"ref_type"`
	MasterBranch string `json:"master_branch"`
	Description  string `json:"description"`
	PusherType   string `json:"pusher_type"`
}

type ActivityMessage struct {
	Message string
}
