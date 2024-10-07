package main

type Post struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	UserId  int    `json:"user_id"`
	Flagged bool   `json:"flagged"`
}
type User struct {
	Id      string
	Name    string
	Role    string
	Blocked bool
}

var users = []User{
	{
		Id:      "1",
		Name:    "Kunal",
		Role:    "admin",
		Blocked: false,
	},
	{
		Id:      "2",
		Name:    "Alice",
		Role:    "member",
		Blocked: false,
	},
}

var posts = []Post{
	{
		Id:      "3123",
		Title:   "Kubernetes API Guide",
		Summary: "Covers the fundamental aspects of K8s api",
		UserId:  2,
		Flagged: false,
	},
	{
		Id:      "4636",
		Title:   "Golang Basics 101",
		Summary: "Covers the basics of Golang",
		UserId:  1,
		Flagged: false,
	},
}
