package main

import (
	"context"
	"log"
	"net/http"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		// extract user id from header
		userId := req.Header.Get("user_id")

		// getting the user role
		var userRole string
		for _, user := range users {
			if userId == user.Id {
				userRole = user.Role
			}
		}

		// cerbos client
		cerbosClient, err := cerbos.New("localhost:3592", cerbos.WithPlaintext())
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		// configure the action based on request
		var action string
		reqMethod := req.Method
		if reqMethod == "GET" {
			action = "read"
		} else if reqMethod == "POST" {
			action = "create"
		} else if reqMethod == "PUT" || reqMethod == "PATCH" {
			action = "update"
		} else if reqMethod == "DELETE" {
			action = "delete"
		}

		// reate a new principle object
		principal := cerbos.NewPrincipal(userId, userRole)

		// create a new resource object
		resource := cerbos.NewResource("posts", action)

		batch := cerbos.NewResourceBatch().Add(resource, action)

		resp, err := cerbosClient.CheckResources(context.Background(), principal, batch)
		if err != nil {
			log.Fatalf("Failed to check resources: %v", err)
		}
		log.Printf("%v", resp)
	})
}
