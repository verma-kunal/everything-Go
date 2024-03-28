package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetAll(w http.ResponseWriter, req *http.Request) {

	// json encoding
	postInBytes, err := json.Marshal(posts)
	if err != nil {
		panic(err.Error())
	}
	w.Write(postInBytes)
}
func GetById(w http.ResponseWriter, req *http.Request) {

	// get id from path
	id := req.PathValue("id")

	for _, post := range posts {
		if id == post.Id {
			toByte, err := json.Marshal(post)
			if err != nil {
				panic(err.Error())
			}
			w.Write(toByte)
		}
	}
}

func main() {

	// register new router
	mux := http.NewServeMux()

	// api routes
	mux.HandleFunc("GET /", GetAll)
	mux.HandleFunc("GET /{id}", GetById)

	server := http.Server{
		Addr:    ":3000",
		Handler: AuthMiddleware(mux), // router wrapped in middleware
	}

	// start server
	log.Print("Listening...")
	server.ListenAndServe()

}
